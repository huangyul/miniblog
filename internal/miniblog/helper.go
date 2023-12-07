package miniblog

import (
	"miniblog/internal/pkg/log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// 定义放置 miniblog 服务配置的默认目录
	recommendedHomeDir = ".miniblog"
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

		// 将home/recommendedHomeDir加入配置文件搜索路径中
		viper.AddConfigPath(filepath.Join(home, recommendedHomeDir))

		// 将当前目录添加到配置文件搜索中
		viper.AddConfigPath(".")

		// 设置配置文件类型
		viper.SetConfigType("yaml")

		// 设置配置文件
		viper.SetConfigFile(defaultConfigName)
	}

	// 读取环境变量
	viper.AutomaticEnv()

	// 读取环境变量前缀为 MINIBLOG，如果是小写，会转为大写
	viper.SetEnvPrefix("MINIBLOG")

	// 将viper.get(key) 中的key中的'.'转为'-'
	replacer := strings.NewReplacer(".", "-")
	viper.EnvKeyReplacer(replacer)

	// 开始读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Errorw("Failed to read viper configuration file", "err", err)
	}

	// 打印viper当前使用的配置文件
	log.Infow("Using config file", "file", viper.ConfigFileUsed())
}

func logOptions() *log.Options {
	return &log.Options{
		DisableCaller:     viper.GetBool("log.disable-caller"),
		DisableStacktrace: viper.GetBool("log.disable-stacktrace"),
		Level:             viper.GetString("log.lever"),
		Format:            viper.GetString("log.format"),
		OutPutPaths:       viper.GetStringSlice("output-paths"),
	}
}
