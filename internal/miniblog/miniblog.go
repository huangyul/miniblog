package miniblog

import (
	"fmt"

	"github.com/spf13/cobra"
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

	return cmd
}

func run() error {
	fmt.Println("hello, miniblog")
	return nil
}
