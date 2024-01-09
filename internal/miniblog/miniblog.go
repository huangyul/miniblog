package miniblog

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewMiniBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "miniblog",
		Short: "a golang practice project",
		Long:  "Long desc Long desc Long desc Long desc Long desc",
		RunE: func(cmd *cobra.Command, args []string) error {
			run()
			return nil
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q dose not take any argument, got %q", cmd.CommandPath(), arg)
				}
			}
			return nil
		},
	}

	return cmd
}

func run() {
	fmt.Println("hello world")
}
