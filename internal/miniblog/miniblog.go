package miniblog

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewMiniblogCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "miniblog",
		Short:        "a go practical project",
		Long:         "long text",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
		Args: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd
}

func run() error {
	fmt.Print("hello, miniblog")

	return nil
}
