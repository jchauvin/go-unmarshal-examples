package models

type Plugin struct {
	URI            string `yaml:"uri"`
	ExpectedResult uint   `yaml:"expectedResult"`
}

type Config struct {
	HostPort     int      `yaml:"hostPort"`
	MetricsPorts int      `yaml:"metricsPort"`
	MetricsPath  string   `yaml:"metricsPath"`
	LogLevel     string   `yaml:"logLevel"`
	Plugins      []Plugin `yaml:"plugins"`
}
