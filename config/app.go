package config

type AppConfig struct {
	Host          string
	Port          string
	IntervalHours int `yaml:"intervalHours"`
}
