package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func Init(mode string) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("config/" + mode + ".yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	_ = viper.BindEnv("ports_grpc", "APP_RPC")
}
