package config

type LoggerConfig struct {
	Level      string `mapstructure:"level" yaml:"level"`
	Dir        string `mapstructure:"dir" yaml:"dir"`
	Filename   string `mapstructure:"filename" yaml:"filename"`
	Format     string `mapstructure:"format" yaml:"format"`
	MaxBackups int    `mapstructure:"max_backups" yaml:"max_backups"`
	MaxDays    int    `mapstructure:"max_days" yaml:"max_days"`
	MaxSize    int    `mapstructure:"max_size" yaml:"max_size"` // unit: MB
	Compress   bool   `mapstructure:"compress" yaml:"compress"`
}
