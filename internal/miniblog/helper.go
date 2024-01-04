package miniblog

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"miniblog/internal/pkg/log"
	"os"
	"path/filepath"
	"strings"
)

const (
	// 默认目录
	recommendedHomeDir = "configs"
	// 指定默认配置文件名
	defaultConfigName = "miniblog.yaml"
)

func initConfig() {
	if cfgFile != "" {
		// 从命令行选项指定的配置文件中读取
		viper.SetConfigFile(cfgFile)
	} else {
		// 查找用户主目录
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// 设置 viper 的搜索路径
		viper.AddConfigPath(filepath.Join(home, recommendedHomeDir))

		viper.AddConfigPath(".")

		// 设置配置文件格式为 yaml
		viper.SetConfigType("yaml")

		// 配置文件名称
		viper.SetConfigName(defaultConfigName)
	}

	// 读取匹配的环境变量
	viper.AutomaticEnv()

	// 读取环境变量的前缀为 MINIBLOG
	viper.SetEnvPrefix("MINIBLOG")

	replacer := strings.NewReplacer(".", "-")
	viper.SetEnvKeyReplacer(replacer)

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
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
