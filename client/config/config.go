package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var App values

type values struct {
	ServerAddress     string
	JsonDataFile      string
	GRPCServerAddress string
}

func init() {
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("client/config")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	if err := viper.Unmarshal(&App); err != nil {
		panic(err)
	}
}
