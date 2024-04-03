package utils

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The value are read by viper from a config file or environment variables.
type Config struct {
	Environment string   `mapstructure:"Environment"`
	ListenIP    string   `mapstructure:"ListenIP"`
	ListenPort  string   `mapstructure:"ListenPort"`
	KeyCloak    KeyCloak `mapstructure:"KeyCloak"`
}

type KeyCloak struct {
	Realm               string  `mapstructure:"Realm"`
	BaseUrl             string  `mapstructure:"BaseUrl"`
	RestApi             RestApi `mapstructure:"RestApi"`
	RealmRS256PublicKey string  `mapstructure:"RealmRS256PublicKey"`
}

type RestApi struct {
	ClientId     string `mapstructure:"ClientId"`
	ClientSecret string `mapstructure:"ClientSecret"`
}

// LoadConfig reads configuration from file or environment variable.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
