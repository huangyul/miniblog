package miniblog

import (
	"miniblog/internal/log"
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

func logOptions() *log.Options {
	return &log.Options{
		DisableCaller:     viper.GetBool("log.disable-caller"),
		DisableStacktrace: viper.GetBool("log.disable-stacktrace"),
		Level:             viper.GetString("log.level"),
		Format:            viper.GetString("log.format"),
		OutputPaths:       viper.GetStringSlice("log.output-paths"),
	}
}
