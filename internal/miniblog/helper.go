package miniblog

import (
	"fmt"
	"miniblog/internal/miniblog/store"
	"miniblog/internal/pkg/log"
	"miniblog/pkg/db"
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

func initStore() error {
	dbOptions := &db.MySqlOptions{
		Host:                  viper.GetString("db.host"),
		Username:              viper.GetString("db.username"),
		Password:              viper.GetString("db.password"),
		Database:              viper.GetString("db.database"),
		MaxIdleConnections:    viper.GetInt("db.max-idle-connections"),
		MaxOpenConnections:    viper.GetInt("db.max-open-connections"),
		MaxConnectionLifeTime: viper.GetDuration("db.max-connection-life-time"),
		LogLevel:              viper.GetInt("db.log-level"),
	}

	ins, err := db.NewMySQL(dbOptions)
	if err != nil {
		return err
	}

	_ = store.NewStore(ins)

	return nil
}
