package miniblog

import (
	"fmt"
	"miniblog/internal/log"
	"net/http"

	"github.com/gin-gonic/gin"
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

	gin.SetMode(viper.GetString("runmode"))

	g := gin.New()

	g.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "not found"})
	})

	g.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	httpSrv := &http.Server{Addr: viper.GetString("addr"), Handler: g}

	log.Infow(fmt.Sprintf(`server is running at %v`, viper.GetString("addr")))

	if err := httpSrv.ListenAndServe(); err != nil {
		log.Errorw(err.Error())
		panic(err)
	}

	return nil
}
