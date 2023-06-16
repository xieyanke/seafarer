package setting

import (
	"time"
)

type ServerSetting struct {
	RunMode           string
	HttpPort          string
	ReadHeaderTimeout time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	MaxHeaderBytes    int
}

type AppSetting struct {
	Name        string
	PageSize    int
	MaxPageSize int
}

type LoggerSetting struct {
	Level      string
	Dir        string
	Filename   string
	MaxSize    string // unit: MB
	MaxBackups int
	MaxAge     int
	Compress   bool
}

type SqliteSetting struct {
	Dir      string
	Filename string
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	return s.vip.UnmarshalKey(k, v)
}
