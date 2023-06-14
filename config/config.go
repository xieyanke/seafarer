package config

type configuration struct {
	App app `mapstructure:"app" yaml:"app"`
}

var Config = new(configuration)
