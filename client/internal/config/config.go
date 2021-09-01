package config

import (
	"fmt"
	"github.com/spf13/viper"
)

//var App values
//
//type values struct {
//	ServerAddress string
//	JsonDataFile  string
//	PortRPC       string
//}
//
//func init() {
//	viper.SetConfigName("config.yaml")
//	viper.SetConfigType("yaml")
//	viper.AddConfigPath("config")
//
//	if err := viper.ReadInConfig(); err != nil {
//		panic(fmt.Errorf("Fatal error config file: %w \n", err))
//	}
//
//	if err := viper.Unmarshal(&App); err != nil {
//		panic(err)
//	}
//}

func Init(mode string) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("config/" + mode + ".yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	viper.SetEnvPrefix("app")
	_ = viper.BindEnv("rpc_server_address", "rpc")
}
