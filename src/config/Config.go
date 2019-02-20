package Config

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Port string
	Mongo struct {
		Name string
		Url string
		Port string
	}
}

func GetConfig() (*Config, error) {
	config, err := InitViper()
	if (err != nil) {
		return &config, err
	}

	return &config, err
}

func InitViper() (Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if (err != nil) {
		return Config{}, err
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.SetDefault("PORT", "8080")
	if err = viper.ReadInConfig(); err != nil {
		log.Panicf("Error reading config file, %s", err)
	}

	var config Config
	err = viper.Unmarshal(&config)
	return config, err
}