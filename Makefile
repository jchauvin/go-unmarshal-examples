
SHELL:=/bin/bash 

export
GOFILES=$(shell find . -type f -name '*.go' -not -path "./vendor/*")
GOPATH=$(shell go env GOPATH)
GIT_COMMIT_SHA=$(shell git rev-parse HEAD 2>/dev/null)
GIT_REMOTE_URL=$(shell git config --get remote.origin.url 2>/dev/null)
BUILD_DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

OS=$(shell uname)
NANCY_VER=1.0.33

.PHONY: all
all: deps lintall gosec opensource scan-nancy tests

.PHONY: pre-commit
pre-commit:
	$(MAKE) deps lintall gosec tests scan-nancy 

.PHONY: lintall
lintall: fmt lint

.PHONY: tests
tests: deps unit coverage

.PHONY: deps
deps:
	go mod download

.PHONY: unit
unit: deps
	go install github.com/onsi/ginkgo/v2/ginkgo@latest
	${GOPATH}/bin/ginkgo -r -p --keep-going -trace -v --randomize-all -race --show-node-events --flake-attempts 3 *

.PHONY: coverage
coverage: deps
	go test ./... -coverprofile cover.out
	go tool cover -func cover.out

.PHONY: fmt
fmt:
	@gofmt -d ${GOFILES}; \
	if [ -n "$$(gofmt -l ${GOFILES})" ]; then \
		echo "Please run 'make dofmt'" && exit 1; \
	fi

.PHONY: dofmt
dofmt:
	go fmt ./...

.PHONY: lint
lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${GOPATH}/bin
	${GOPATH}/bin/golangci-lint run ./...

.PHONY: pre-push
pre-push:
	golangci-lint run ./...
	$(MAKE) fmt
	go test ./...

.PHONY: gosec
gosec:
	curl -sSfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b ${GOPATH}/bin v2.15.0
	${GOPATH}/bin/gosec -quiet --exclude=G104 ./...

# .PHONY: buildgo-linux
# buildgo-linux: deps
# 	CGO_ENABLED=0 GOOS=linux go build -ldflags '-s -w -extldflags "-static"' -o cmd/logger-external-exporter cmd/main.go

.PHONY: scan-nancy
scan-nancy:
	if [[ "$(shell nancy -V | grep -o 'nancy version')" != "nancy version" ]]; then \
	   if [[ "${OS}" == "Darwin" ]]; then \
		   curl -Lo /tmp/nancy.tar.gz https://github.com/sonatype-nexus-community/nancy/releases/download/v${NANCY_VER}/nancy-v${NANCY_VER}-darwin-amd64.tar.gz; \
	   else \
		   curl -Lo /tmp/nancy.tar.gz https://github.com/sonatype-nexus-community/nancy/releases/download/v${NANCY_VER}/nancy-v${NANCY_VER}-linux-amd64.tar.gz; \
	   fi; \
	   sudo tar -xzvf /tmp/nancy.tar.gz -C /usr/local/bin/ nancy; \
	fi
	go list -json -m all | nancy sleuth --skip-update-check
