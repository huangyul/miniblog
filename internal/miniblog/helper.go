package miniblog

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// 配置文件名
	defaultConfigName = "miniblog.yaml"
)

// initConfig 设置需要读取的配置文件名，环境变量，并读取配置文件的内容到viper中
func initConfig() {
	if cfgFile != "" {
		// 从命令行选项中指定的文件读取
		viper.SetConfigFile(cfgFile)
	} else {
		// 获取用户主目录
		home, err := os.UserHomeDir()
		// 如果获取失败，则打印错误信息，并退出程序
		cobra.CheckErr(err)

		// 将获取的主目录添加到配置文件搜索中
		viper.AddConfigPath(home)

		// 设置配置文件类型
		viper.SetConfigType("yaml")

		// 设置配置文件
		viper.SetConfigFile(defaultConfigName)
	}

	// 读取环境变量
	viper.AutomaticEnv()

	// 开始读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
