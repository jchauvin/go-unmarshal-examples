package models

import "gopkg.in/yaml.v3"

type Plugin struct {
	Type   string    `yaml:"type"`
	Config yaml.Node `yaml:"config"`
}

type Config struct {
	HostPort     int      `yaml:"hostPort"`
	MetricsPorts int      `yaml:"metricsPort"`
	MetricsPath  string   `yaml:"metricsPath"`
	LogLevel     string   `yaml:"logLevel"`
	Plugins      []Plugin `yaml:"plugins"`
}

type HTTPPluginConfig struct {
	URI            string `yaml:"uri"`
	ExpectedResult uint   `yaml:"expectedResult"`
}

type SMPluginConfig struct {
	URI         string   `yaml:"uri"`
	IAMAuthType string   `yaml:"iamAuthType"`
	Secrets     []Secret `yaml:"secrets"`
}

type Secret struct {
	SecretID   string `yaml:"secretID"`
	SecretName string `yaml:"secretName"`
}

type PluginConfig interface {
	Validate() error
}
