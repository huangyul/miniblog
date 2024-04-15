package miniblog

import (
	"encoding/json"
	"fmt"

	"github.com/huangyul/miniblog/internal/pkg/log"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var cfgFile = ""

func NewMiniBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		// 该命令的名称
		Use:          "miniblog",
		Short:        "miniblog server1",
		Long:         "miniblog server2",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Init(log.NewOptions())
			defer log.Sync()

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

	cobra.OnInitialize(initConfig)

	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "The path to the miniblog configuration file. Empty string for no configuration file.")

	cmd.Flags().BoolP("toggle", "t", false, "help message for toggle")

	return cmd
}

func run() error {
	settings, _ := json.Marshal(viper.AllSettings())
	log.Infow(string(settings))

	log.Infow(viper.GetString("db.username"))
	return nil
}
