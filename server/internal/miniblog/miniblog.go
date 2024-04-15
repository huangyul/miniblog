package miniblog

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewMiniBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		// 该命令的名称
		Use:          "miniblog",
		Short:        "miniblog server1",
		Long:         "miniblog server2",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
		Args: func(cmd *cobra.Command, args []string) error {
			// 命令不接受任何参数
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any argments, got %q", cmd.CommandPath(), arg)
				}
			}
			return nil
		},
	}

	return cmd
}

func run() error {
	fmt.Println("miniblog server")
	return nil
}
