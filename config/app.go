package config

import (
	"fmt"
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	LogLevel    string `yaml:"logLevel" yaml_default:"info"`
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
	AutoDownloadDir string `yaml:"autoDownloadDir"`
	Transmission    struct {
		RpcUrl   string `yaml:"rpcUrl"`
		Login    string
		Password string
		Folders  map[string]string
	}
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
