package miniblog

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/huangyul/miniblog/internal/pkg/core"
	"github.com/huangyul/miniblog/internal/pkg/errno"
	"github.com/huangyul/miniblog/internal/pkg/log"
	"github.com/huangyul/miniblog/internal/pkg/middleware"
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

	gin.SetMode(viper.GetString("runmode"))

	server := gin.Default()

	server.Use(middleware.Cors(), middleware.RequestID())

	server.NoRoute(func(ctx *gin.Context) {
		core.WriteResponse(ctx, errno.ErrPageNotFound, nil)
	})
	server.GET("/healthz", func(ctx *gin.Context) {
		core.WriteResponse(ctx, nil, map[string]string{"status": "ok"})
	})

	httpSrv := &http.Server{
		Addr:    viper.GetString("addr"),
		Handler: server,
	}

	go func() {
		if err := httpSrv.ListenAndServe(); err != nil {
			log.Fatalw(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Infow("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := httpSrv.Shutdown(ctx); err != nil {
		log.Errorw("Insecure Server forced to shutdown", "err", err)
		return err
	}

	return nil
}
