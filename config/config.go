package config

import (
	"sync"

	"github.com/spf13/viper"
)

var once sync.Once

type Config struct {
	Url      string `mapstructure:"SCRAPING_URL"`
	LogLevel string `mapstructure:"LOG_LEVEL"`
}

var Conf *Config = &Config{}

func GetConfig() *Config {
	once.Do(func() {
		viper.SetDefault("SCRAPING_URL", "https://candystore.zimpler.net/#candystore-customers")
		viper.SetDefault("LOG_LEVEL", "INFO")
		viper.SetConfigFile(".env")
		viper.ReadInConfig()
		viper.Unmarshal(Conf)
	})
	return Conf
}
