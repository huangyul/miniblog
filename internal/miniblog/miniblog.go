package miniblog

import (
	"encoding/json"
	"fmt"
	"miniblog/internal/log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
			log := log.NewLogger(logOptions())
			defer log.Sync()

			return run()
		},
	}

	cobra.OnInitialize(initConfig)

	return cmd
}

func run() error {
	fmt.Println("miniblog")
	setting, _ := json.Marshal(viper.AllSettings())
	log.Infow(string(setting))
	log.Infow(viper.GetString("db.username"))
	return nil
}
