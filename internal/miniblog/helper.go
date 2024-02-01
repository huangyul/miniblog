package miniblog

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
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
