package config

type configuration struct {
	App    appConfig    `mapstructure:"app" yaml:"app"`
	Logger LoggerConfig `mapstructure:"logger" yaml:"logger"`
}

var Config = new(configuration)
