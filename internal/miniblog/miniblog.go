package miniblog

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewMiniBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "miniblog",
		Short:        "miniblog serve",
		Long:         "miniblog long detail",
		SilenceUsage: true,

		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("command %q not found", arg)
				}
			}
			return nil
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}

	return cmd
}

func run() error {
	fmt.Println("miniblog")
	return nil
}
