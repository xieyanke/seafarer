package setting

import (
	"github.com/spf13/viper"
)

type Setting struct {
	vip *viper.Viper
}

func NewSetting(fileDir, fileName, fileType string) (*Setting, error) {
	vip := viper.New()

	vip.AddConfigPath(fileDir)
	vip.SetConfigName(fileName)
	vip.SetConfigType(fileType)
	err := vip.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vip}, nil
}
