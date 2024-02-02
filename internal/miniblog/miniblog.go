package miniblog

import (
	"encoding/json"
	"miniblog/internal/pkg/log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewMiniblogCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "miniblog",
		Short:        "a go practical project",
		Long:         "long text",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Init(logOptions())
			defer log.Sync()

			return run()
		},
		Args: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	cobra.OnInitialize(initConfig)

	return cmd
}

func run() error {
	settings, _ := json.Marshal(viper.AllSettings())
	log.Infow(string(settings))

	log.Infow(viper.GetString("db.username"))

	return nil
}
