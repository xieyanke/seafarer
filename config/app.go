package config

type app struct {
	Name string `mapstructure:"name" yaml:"name"`
	Port string `mapstructure:"port" yaml:"port"`
	Env  string `mapstructure:"env" yaml:"env"`
}
