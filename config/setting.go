package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Data Config

type AppEnv struct {
	Environment   string `mapstructure:"environment"`
	JWT           string `mapstructure:"jwtkey"`
	Port          string `mapstructure:"port"`
	AllowedOrigin string `mapstructure:"allowedorigin"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Name     string `mapstructure:"dbname"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type DefaultSetting struct {
	AdminUsername string `mapstructure:"adminusername"`
	AdminPassword string `mapstructure:"adminpassword"`
}

type Config struct {
	DB      Database       `mapstructure:"database"`
	App     AppEnv         `mapstructure:"app"`
	Default DefaultSetting `mapstructure:"default"`
}

func Init(filename string) {
	viper.SetConfigFile(filename)
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err := viper.Unmarshal(&Data); err != nil {
		panic(fmt.Errorf("unable to decode into struct: %s \n", err))
	}
}
