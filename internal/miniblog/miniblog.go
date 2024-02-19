package miniblog

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewMiniBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "miniblog",
		Short: "Short text",
		RunE: func(cmd *cobra.Command, args []string) error {
			// main.exe -n=test
			name, _ := cmd.Flags().GetString("name")
			if len(name) > 0 {
				fmt.Println("name is ", name)
			}

			// 读取配置文件
			cfg, _ := cmd.Flags().GetString("config")
			if len(cfg) > 0 {
				configFile = cfg
			}
			initConfig()

			return run()
		},
		SilenceUsage: true,
		// 额外的参数，例如 main.exe  xxx
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("command: %v is not exits", arg)
				}
			}
			return nil
		},
	}

	cmd.Flags().StringP("name", "n", "", "--n")
	cmd.Flags().StringP("config", "c", "", "配置文件路径")

	return cmd
}

func run() error {
	fmt.Println(viper.GetString("db.username"))

	fmt.Println("hello, miniblog")

	return nil
}
