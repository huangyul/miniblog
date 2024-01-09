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

	cobra.OnInitialize(initConfig)

	return cmd
}

func run() {
	data, _ := json.Marshal(viper.AllSettings())
	fmt.Println(string(data))
	fmt.Println(viper.GetString("db.username"))
}
