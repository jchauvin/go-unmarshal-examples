package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/chauvinj/unmarshal-example/json/pkg/config"
	"github.com/chauvinj/unmarshal-example/json/pkg/models"
	//"github.com/davecgh/go-spew/spew"
	"github.com/fatih/color"
)

var (
	CommitSHA      string = "no-version-info"
	configFilePath        = flag.String("conf", "config.json", "The config path file")
)

func main() {

	var (
		err error
		cfg *models.Config
	)

	color.Green("Unmarshalling delayed json example")

	cfg, err = setup()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read config - error: %v", err)
		os.Exit(-1)
	}

	c := color.New(color.FgYellow).Add(color.Bold)
	p := color.New(color.FgHiGreen)
	c.Printf("============================================\n")
	c.Printf("Full config ( with raw message )- %+v\n", cfg)
	c.Printf("----------------------------------\n")
	p.Printf("\nPlugins -\n")
	for _, plugin := range cfg.Plugins {
		p.Printf("----------------------------------\n")
		p.Printf("\nPlugin of Type - %s\n", plugin.Type)
		p.Printf("\tConfig Raw - %+v\n", plugin.Config)
		config.UnmarshalPlugin(&plugin)
		p.Printf("----------------------------------\n")
	}
	c.Printf("----------------------------------\n\n")
	str, _ := json.MarshalIndent(cfg, "", "\t")
	c.Printf("Json Marshalled - ")
	c.Println(string(str))
	c.Printf("\n============================================\n")

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
