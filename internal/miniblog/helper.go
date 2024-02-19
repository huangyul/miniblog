package miniblog

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	configFile = ""
)

func initConfig() {
	if len(configFile) != 0 {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath("./config")
		viper.SetConfigName("miniblog.yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("use config file is", viper.ConfigFileUsed())
}
