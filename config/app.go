package config

import (
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	Host        string
	Port        string
	Interval    time.Duration `yaml:"interval"`
	Credentials map[string]struct {
		Login    string
		Password string
	}
	Notifications struct {
		TrayMessage   bool `yaml:"trayMessage"`
		KDEMessage    bool `yaml:"KDEMessage"`
		OpenInBrowser bool `yaml:"openInBrowser"`
		OpenFile      bool `yaml:"openFile"`
	}
}

func Load(filePath string) *AppConfig {
	config := &AppConfig{}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	err = yaml.Unmarshal(data, config)
	if err != nil {
		log.Fatalf("error parsing config: %v", err)
	}

	return config
}
