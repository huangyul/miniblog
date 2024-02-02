package miniblog

import (
	"fmt"
	"miniblog/internal/pkg/log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

var (
	defaultConfigName = "miniblog.yaml"
)

func initConfig() {
	viper.AddConfigPath(filepath.Join("."))
	viper.SetConfigType("yaml")
	viper.SetConfigName(defaultConfigName)

	viper.AutomaticEnv()

	replacer := strings.NewReplacer(".", "-")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Fprintln(os.Stdout, "Using config file:", viper.ConfigFileUsed())
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
