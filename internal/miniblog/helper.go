package miniblog

import (
	"strings"

	"github.com/spf13/viper"
)

var (
	configName = "miniblog.yaml"
)

func initConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName(configName)

	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.EnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
