package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config interface {
	GetString(key string) string
	GetInt(key string) int
	GetBool(key string) bool
}

type viperConfig struct{}

var profile string

func (v *viperConfig) Init() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("/opt/app/")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic("No file config in main directory")
	}

	profile = os.Getenv("MODE")
	if profile == "" {
		profile = "development"
	}
}

func (v *viperConfig) GetString(key string) string {
	return viper.Sub(profile).GetString(key)
}

func (v *viperConfig) GetInt(key string) int {
	return viper.Sub(profile).GetInt(key)
}

func (v *viperConfig) GetBool(key string) bool {
	return viper.Sub(profile).GetBool(key)
}

func NewViperConfig() Config {
	v := &viperConfig{}
	v.Init()
	return v
}
