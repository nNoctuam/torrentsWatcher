package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type AppConfig struct {
	Host          string
	Port          string
	IntervalHours int `yaml:"intervalHours"`
	Credentials   map[string]struct {
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

func Load() *AppConfig {
	config := &AppConfig{}

	dat, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	err = yaml.Unmarshal(dat, config)
	if err != nil {
		log.Fatalf("error parsing config: %v", err)
	}

	return config
}
