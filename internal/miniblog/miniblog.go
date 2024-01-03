package miniblog

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func NewMiniBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		// 指定命令的名字
		Use: "miniblog",
		// 命令的简短描述
		Short: "A good Go practical project",
		// 命令的详细描述
		Long: `A good Go practical project, used to create user with basic information.

Find more miniblog information at:
	https://github.com/huangyul/miniblog#readme`,
		// 命令出错时，不打印帮助信息；方便看到错误信息
		SilenceUsage: true,
		// 指定调用 cmd.Execute() 时，执行Run函数
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
		// 这里设置命令运行时，不需要指定命令行参数
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q dose not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}

			return nil
		},
	}

	// 使 initconfig 函数在每个命令运行时都会被调用以读取配置
	cobra.OnInitialize(initConfig)

	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "The path to the miniblog configuration file, Empty string fro no configuration file.")

	// 本地标志
	cmd.Flags().BoolP("toogle", "t", false, "help message for toggle")

	return cmd
}

// run 是实际业务代码入口函数
func run() error {
	setting, _ := json.Marshal(viper.AllSettings())
	fmt.Println(string(setting))
	fmt.Println(viper.GetString("db.username"))
	return nil
}
