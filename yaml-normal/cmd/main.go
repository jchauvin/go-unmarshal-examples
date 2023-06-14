package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/chauvinj/unmarshal-example/yaml-normal/pkg/config"
	"github.com/chauvinj/unmarshal-example/yaml-normal/pkg/models"
	"github.com/fatih/color"
	"gopkg.in/yaml.v3"
)

var (
	CommitSHA      string = "no-version-info"
	configFilePath        = flag.String("conf", "config.yaml", "The config path file")
)

func main() {

	var (
		err error
		cfg *models.Config
	)

	color.Green("Unmarshalling delayed yaml V3 example")

	cfg, err = setup()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read config - error: %v", err)
		os.Exit(-1)
	}

	c := color.New(color.FgYellow).Add(color.Bold)
	p := color.New(color.FgHiGreen)
	c.Println("============================================")
	c.Printf("Full config - %+v\n\n", cfg)
	c.Printf("----------------------------------\n")
	p.Printf("\nPlugins -\n")
	for _, plugin := range cfg.Plugins {
		p.Printf("----------------------------------\n")
		p.Printf("Plugin Config  - %#v\n", plugin)
		yamlData, _ := yaml.Marshal(plugin)
		p.Printf("Plugin Marshalled - \n")
		p.Println(string(yamlData))
		p.Printf("----------------------------------\n")
	}
	c.Printf("----------------------------------\n\n")

	yamlData, err := yaml.Marshal(cfg)
	if err != nil {
		log.Fatalf("Failed to marshal YAML: %v", err)
	}
	c.Printf("Full config Yaml Marshalled - \n")
	c.Println(string(yamlData))
	c.Printf("\n============================================")

}

func setup() (*models.Config, error) {
	// Parse config path
	flag.Parse()

	cfg, err := config.LoadConfig(*configFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load config - %w", err)
	}

	return cfg, nil

}
