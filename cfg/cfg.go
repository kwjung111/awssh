package cfg

import (
	"fmt"
	"log"
	"sync"

	"github.com/spf13/viper"
)

// define default value
type Config struct {
	Profile string
	KeyFile string
	Port    int
}

var once sync.Once

var conf Config

func InitCfgManager() *Config {

	//TODO modify path
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {

		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found, creating default config")

			// defualt setting
			viper.Set("profile", "")
			viper.Set("keyFile", "key.pem")
			viper.Set("port", 22)

			if err := viper.SafeWriteConfig(); err != nil {
				log.Fatalf("Error creating config file, %s", err)
			}
		} else {
			log.Fatalf("Error reading config file, %s", err)
		}
	}

	conf.Profile = viper.GetString("profile")
	conf.KeyFile = viper.GetString("keyFile")
	conf.Port = viper.GetInt("port")

	return &conf
}

func GetConf() *Config {
	once.Do(func() {
		InitCfgManager()
	})
	return &conf
}

func SaveItem(vip *viper.Viper, name string) error {

	ConfigureViperDefaults(vip)
	err := vip.WriteConfigAs(name)
	if err != nil {
		return fmt.Errorf("failed to save configuration: %w", err)
	}

	fmt.Printf("config saved : %s\n", name)
	return nil
}

func LoadItem(name string) *viper.Viper {

	vip := viper.New()
	ConfigureViperDefaults(vip)
	vip.SetConfigFile(name) //no extension

	return vip
}

func ConfigureViperDefaults(vip *viper.Viper) {
	vip.AddConfigPath(".")
	vip.SetConfigType("yaml")
}
