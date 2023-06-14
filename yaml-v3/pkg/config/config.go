package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/chauvinj/unmarshal-example/yaml-v3/pkg/models"
	"gopkg.in/yaml.v3"
)

// LoadConfig loads the config from the file specified in params
func LoadConfig(path string) (*models.Config, error) {

	var (
		cfg   models.Config
		bytes []byte
		err   error
	)

	//bytes, err = os.ReadFile(filepath.Join("/", filepath.Clean(path)))
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

	t := p.Type
	switch t {
	case "http":
		c := &models.HTTPPluginConfig{}
		_ = p.Config.Decode(c)
		fmt.Printf("\nPlugin config unmarshalled struct - \n\t%#v\n", c)
		str, _ := yaml.Marshal(c)
		fmt.Printf("\nPlugin config Marshalled - \n")
		fmt.Printf("\t%s\n", string(str))
	case "secrets_manager":
		c := &models.SMPluginConfig{}
		_ = p.Config.Decode(c)
		fmt.Printf("\nPlugin config unmarshalled struct - \n\t%#v\n", c)
		str, _ := yaml.Marshal(c)
		fmt.Printf("\nPlugin config Marshalled - \n")
		fmt.Printf("%s\n", string(str))
	default:
		return fmt.Errorf("unkonwn plugin type - %s", t)
	}

	return nil

}
