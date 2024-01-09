package miniblog

import (
	"encoding/json"
	"fmt"
	"miniblog/internal/log"

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
			log.Init(logOptions())
			defer log.Sync()

			return run()
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

func run() error {
	data, _ := json.Marshal(viper.AllSettings())
	log.Infow(string(data))
	log.Infow(viper.GetString("db.username"))

	return nil
}
