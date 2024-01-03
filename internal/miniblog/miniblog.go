package miniblog

import (
	"fmt"
	"github.com/spf13/cobra"
)

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

	return cmd
}

// run 是实际业务代码入口函数
func run() error {
	fmt.Print("hello miniblog")
	return nil
}
