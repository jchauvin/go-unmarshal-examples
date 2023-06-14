package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/chauvinj/unmarshal-example/yaml-v2/pkg/models"
	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
)

// LoadConfig loads the config from the file specified in params
func LoadConfig(path string) (*models.Config, error) {

	var (
		cfg   models.Config
		bytes []byte
		err   error
	)

	bytes, err = os.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s with error %v", path, err)
	}

	err = yaml.Unmarshal(bytes, &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal object %s with error %v", bytes, err)
	}

	err = validateConfig(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func validateConfig(cfg *models.Config) error {

	validator := models.Validator{}

	validator.AddConstraint(models.PositiveNotZero(cfg.MetricsPorts, "MetricsPort"))
	validator.AddConstraint(models.PositiveNotZero(cfg.HostPort, "HostPort"))
	validator.AddConstraint(models.NotEmpty(cfg.MetricsPath, "MetricsPath"))
	validator.AddConstraint(models.NotZeroLength(cfg.Plugins, "Plugins"))

	return validator.Validate()
}

func UnmarshalPlugin(p *models.Plugin) error {

	cp := color.New(color.FgHiBlue).Add(color.Bold)
	t := p.Type
	switch t {
	case "http":
		c := &models.HTTPPluginConfig{}
		p.Config.Unmarshal(c)
		cp.Printf("\nPlugin config unmarshalled struct - \n\t%#v\n", c)
		str, _ := yaml.Marshal(c)
		cp.Printf("\nPlugin config Marshalled - \n")
		cp.Printf("\t%s\n", string(str))
	case "secrets_manager":
		c := &models.SMPluginConfig{}
		p.Config.Unmarshal(c)
		cp.Printf("\nPlugin config unmarshalled struct - \n\t%#v\n", c)
		str, _ := yaml.Marshal(c)
		cp.Printf("\nPlugin config Marshalled - \n")
		cp.Printf("%s\n", string(str))
	default:
		return fmt.Errorf("unkonwn plugin type - %s", t)
	}

	return nil
	
}
