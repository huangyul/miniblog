package miniblog

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"miniblog/internal/pkg/log"

	"github.com/spf13/cobra"
)

var cfgFile string

// NewMiniBlogCommand 构建一个 cobra.Command 对象，之后可以使用 Execute 来启动程序
func NewMiniBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		// 指定命令的名字
		Use: "miniblog",
		// 命令的简短描述
		Short: "A Go practical project",
		// 命令的详细描述
		Long: `A Go practical project, used to create user with basic information.

Find more miniblog information at:
		https://github.com/huangyul/miniblog#readme`,
		// 命令出错时，不打印帮助信息。设置为true可以保持命令出错时一眼看到错误信息
		SilenceUsage: true,
		// 指定调用cmd.Execute()时，执行Run函数，函数执行失败会返回错误信息
		RunE: func(cmd *cobra.Command, args []string) error {
			// 初始化日志
			log.Init(logOptions())
			defer log.Sync() // Sync 将缓存中的日志刷新到磁盘中
			return run()
		},
		// 这里设置命令运行时，不需要指定命令行参数
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}

	// 使initConfig函数在每个命令执行都会被调用
	cobra.OnInitialize(initConfig)

	// Cobra 支持持久性标志(PersistentFlag)，该标志可用于它所分配的命令以及该命令下的每个子命令
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "The path to the miniblog configuration file. Empty string for no configuration file.")

	// Cobra 也支持本地标志，本地标志只能在其所绑定的命令上使用
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return cmd
}

// 实际业务代码入口函数
func run() error {
	// 获取viper读到的所有配置项
	settings, _ := json.Marshal(viper.AllSettings())
	log.Infow(string(settings))

	// 打印db -> username 配置的值
	fmt.Println(viper.GetString("db.username"))
	log.Infow(string(viper.GetString("db.username")))

	return nil
}
