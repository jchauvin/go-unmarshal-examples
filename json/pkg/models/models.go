package models

import (
	"encoding/json"
)

type Plugin struct {
	Type   string          `json:"type"`
	Config json.RawMessage `json:"config"`
}

type Config struct {
	HostPort     int      `json:"hostPort"`
	MetricsPorts int      `json:"metricsPort"`
	MetricsPath  string   `json:"metricsPath"`
	LogLevel     string   `json:"logLevel"`
	Plugins      []Plugin `json:"plugins"`
}

type HTTPPluginConfig struct {
	URI            string `json:"uri"`
	ExpectedResult uint   `json:"expectedResult"`
}

type SMPluginConfig struct {
	URI         string   `json:"uri"`
	IAMAuthType string   `json:"iamAuthType"`
	Secrets     []Secret `json:"secrets"`
}

type Secret struct {
	SecretID   string `json:"secretID"`
	SecretName string `json:"secretName"`
}

type PluginConfig interface {
	Validate() error
}
