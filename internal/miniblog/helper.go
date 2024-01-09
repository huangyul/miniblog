package miniblog

import (
	"fmt"
	"miniblog/internal/log"
	"os"

	"github.com/spf13/viper"
)

var (
	// recommendedHomeDir = "_output"
	defaultConfigName = "miniblog.yaml"
)

func initConfig() {
	if len(cfgFile) > 0 {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(defaultConfigName)
	}

	// 获取环境变量
	viper.AutomaticEnv()
	viper.SetEnvPrefix("miniblog")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, err, "os error")
	}

	fmt.Fprintln(os.Stdout, "using config file:", viper.ConfigFileUsed())
}

// logOptions 从 viper 中读取日志配置
func logOptions() *log.Options {
	return &log.Options{
		DisableCaller:     viper.GetBool("log.disable-caller"),
		DisableStacktrace: viper.GetBool("log.disable-stacktrace"),
		Level:             viper.GetString("log.level"),
		Format:            viper.GetString("log.format"),
		OutputPaths:       viper.GetStringSlice("log.output-paths"),
	}
}
