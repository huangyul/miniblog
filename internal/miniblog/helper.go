package miniblog

import (
	"fmt"
	"miniblog/internal/pkg/log"

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

func logOptions() *log.Options {
	return &log.Options{
		DisableCaller:     viper.GetBool("log.disbable-caller"),
		DisableStacktrace: viper.GetBool("log.disable-stacktrace"),
		Level:             viper.GetString("log.level"),
		Format:            viper.GetString("log.format"),
		OutputPaths:       viper.GetStringSlice("log.output-paths"),
	}
}
