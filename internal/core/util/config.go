package util

import (
	"github.com/spf13/viper"
	"log"
	"monero-shop-api/internal/exception"
)

type Impl interface {
	Get() Config
}

type configImpl struct {
}

func New() Impl {
	return &configImpl{}
}

func (c *configImpl) Get() Config {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.AutomaticEnv()

	var cfg Config

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("Unable to read config: %v", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			exception.PanicLogging(err)
		}
	}

	err = v.Unmarshal(&cfg)
	exception.PanicLogging(err)
	return cfg
}
