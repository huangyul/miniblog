package miniblog

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const (
	defaultConfigName = "miniblog.yaml"
)

func initConfig() {

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")

		viper.SetConfigType("yaml")

		viper.SetConfigName(defaultConfigName)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Fprintln(os.Stdout, "Using config file:", viper.ConfigFileUsed())
}
