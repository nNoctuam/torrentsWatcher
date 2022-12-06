package config

import (
	"fmt"
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

type NotificationsConfig struct {
	OpenInBrowser bool `yaml:"openInBrowser"`
	OpenFile      bool `yaml:"openFile"`
	Message       map[string]bool
}

type AppConfig struct {
	LogLevel    string `yaml:"logLevel" yaml_default:"info"`
	Host        string
	Port        string
	Interval    time.Duration `yaml:"interval"`
	Credentials map[string]struct {
		Login    string
		Password string
	}
	Notifications   NotificationsConfig
	AutoDownloadDir string `yaml:"autoDownloadDir"`
	Transmission    struct {
		RPCURL   string `yaml:"rpcUrl"`
		Login    string
		Password string
		Folders  map[string]string
	}
	BlockViewList []string `yaml:"blockViewList"`
}

func Load(filePath string) (*AppConfig, error) {
	config := &AppConfig{}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("loading config: %w", err)
	}

	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, fmt.Errorf("parsing config: %w", err)
	}

	return config, nil
}
