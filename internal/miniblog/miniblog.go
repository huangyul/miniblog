package miniblog

import (
	"fmt"
	"miniblog/internal/pkg/log"
	"miniblog/internal/pkg/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewMiniBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "miniblog",
		Short: "Short text",
		RunE: func(cmd *cobra.Command, args []string) error {
			// main.exe -n=test
			name, _ := cmd.Flags().GetString("name")
			if len(name) > 0 {
				fmt.Println("name is ", name)
			}

			// 读取配置文件
			cfg, _ := cmd.Flags().GetString("config")
			if len(cfg) > 0 {
				configFile = cfg
			}

			// 初始化 log
			log.Init(logOptions())
			defer log.Sync()

			return run()
		},
		SilenceUsage: true,
		// 额外的参数，例如 main.exe  xxx
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("command: %v is not exits", arg)
				}
			}
			return nil
		},
	}

	cobra.OnInitialize(initConfig)

	cmd.Flags().StringP("name", "n", "", "--n")
	cmd.Flags().StringP("config", "c", "", "配置文件路径")

	return cmd
}

func run() error {

	r := gin.Default()

	r.Use(middleware.RequestID())

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": "not found",
		})
		log.Infow("apt not found", "api:", ctx.Request.URL.Path)
	})

	r.GET("/healthz", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
		log.Infow("healthz api")
	})

	log.Infow("server is running", "addr", viper.GetString("server.addr"))
	err := r.Run(viper.GetString("server.addr"))

	return err
}
