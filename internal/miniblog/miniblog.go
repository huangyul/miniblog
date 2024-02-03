package miniblog

import (
	"context"
	"errors"
	"miniblog/internal/pkg/log"
	"miniblog/internal/pkg/middleware"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
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
	gin.SetMode("debug")

	g := gin.New()

	wm := []gin.HandlerFunc{middleware.Cors(), middleware.NoCache(), middleware.RequestId()}

	g.Use(wm...)

	g.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": "not found",
		})
	})

	g.GET("/healthz", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "api is ok")
	})

	addr := viper.GetString("addr")

	httpSrv := &http.Server{
		Addr:    addr,
		Handler: g,
	}

	log.Infow("server is running at %s", "addr", addr)

	go func() {
		if err := httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalw(err.Error())
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Infow("sutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := httpSrv.Shutdown(ctx); err != nil {
		log.Errorw("Insecure Server forced to shutdown", "err", err)
		return err
	}

	log.Infow("server shutdown success")

	return nil
}
