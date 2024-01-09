package miniblog

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"miniblog/internal/pkg/log"
	"miniblog/internal/pkg/middleware"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	// 使用中间件
	// gin.Recovery() 用来铺货任何 panic 错误，并恢复
	mws := []gin.HandlerFunc{gin.Recovery(), middleware.RequestID(), middleware.NoCache, middleware.Cors, middleware.Secure}
	g.Use(mws...)

	g.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "page not found"})
	})

	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	httpsrv := &http.Server{Addr: viper.GetString("addr"), Handler: g}

	log.Infow("start to listening at", "addr", viper.GetString("addr"))

	go func() {
		if err := httpsrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalw(err.Error())
		}
	}()

	// 等待中断信号优雅关闭服务器（10秒超时）
	quit := make(chan os.Signal)
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的 CTRL + C 就是触发系统 SIGINT 信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// 这里不会阻塞
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 这里会阻塞
	<-quit
	// 当收到信号
	log.Infow("Shutting down server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpsrv.Shutdown(ctx); err != nil {
		log.Errorw("Insecure Server forced to shutdown", "err", err)
		return err
	}

	log.Infow("Server exiting")

	return nil

}
