package miniblog

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"miniblog/internal/log"
	"net/http"
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
	gin.SetMode(viper.GetString("runmode"))

	g := gin.New()

	g.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "page not found"})
	})

	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	httpsrv := &http.Server{Addr: viper.GetString("addr"), Handler: g}

	log.Infow("start to listening at", "addr", viper.GetString("addr"))

	if err := httpsrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalw(err.Error())
	}

	return nil

	return nil
}
