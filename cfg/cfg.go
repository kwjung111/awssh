package cfg

import (
	"fmt"
	"log"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Profile string
}

var once sync.Once

var conf Config

func InitCfgManager() *Config {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {

		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found, creating default config")

			// defualt setting
			viper.Set("profile", "")

			if err := viper.SafeWriteConfig(); err != nil {
				log.Fatalf("Error creating config file, %s", err)
			}
		} else {
			log.Fatalf("Error reading config file, %s", err)
		}
	}

	conf.Profile = viper.GetString("profile")

	return &conf
}

func GetConf() *Config {
	once.Do(func() {
		InitCfgManager()
	})
	return &conf
}
