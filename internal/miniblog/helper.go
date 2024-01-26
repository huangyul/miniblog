package miniblog

import "github.com/spf13/viper"

var (
	configName = "miniblog.yaml"
)

func initConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName(configName)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
