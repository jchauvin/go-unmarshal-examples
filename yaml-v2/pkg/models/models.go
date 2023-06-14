package models

type Plugin struct {
	Type   string      `yaml:"type"`
	Config *RawMessage `yaml:"config"`
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

type RawMessage struct {
	unmarshal func(interface{}) error
}

func (msg *RawMessage) UnmarshalYAML(unmarshal func(interface{}) error) error {
	msg.unmarshal = unmarshal
	return nil
}

func (msg *RawMessage) Unmarshal(v interface{}) error {
	return msg.unmarshal(v)
}

func (p Plugin) MarshalYAML() (interface{}, error) {

	var err error

	t := p.Type
	yamlData := make(map[string]interface{})
	switch t {
	case "http":
		c := &HTTPPluginConfig{}
		err = p.Config.Unmarshal(c)
		yamlData["Type"] = p.Type
		yamlData["Config"] = c
	case "secrets_manager":
		c := &SMPluginConfig{}
		err = p.Config.Unmarshal(c)
		yamlData["Type"] = p.Type
		yamlData["Config"] = c
	}

	if err != nil {
		return yamlData, err
	}

	return yamlData, nil
}
