package miniblog

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"miniblog/internal/pkg/log"
	"net/http"
)

var cfgFile string

func NewMiniBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		// 指定命令的名字
		Use: "miniblog",
		// 命令的简短描述
		Short: "A good Go practical project",
		// 命令的详细描述
		Long: `A good Go practical project, used to create user with basic information.

Find more miniblog information at:
	https://github.com/huangyul/miniblog#readme`,
		// 命令出错时，不打印帮助信息；方便看到错误信息
		SilenceUsage: true,
		// 指定调用 cmd.Execute() 时，执行Run函数
		RunE: func(cmd *cobra.Command, args []string) error {

			// 初始化日志
			log.Init(logOptions())
			defer log.Sync() // Sync 将缓存中的日志刷新到磁盘文件中

			return run()
		},
		// 这里设置命令运行时，不需要指定命令行参数
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q dose not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}

			return nil
		},
	}

	// 使 initconfig 函数在每个命令运行时都会被调用以读取配置
	cobra.OnInitialize(initConfig)

	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "The path to the miniblog configuration file, Empty string fro no configuration file.")

	// 本地标志
	cmd.Flags().BoolP("toogle", "t", false, "help message for toggle")

	return cmd
}

// run 是实际业务代码入口函数
func run() error {

	// 设置 gin 模式
	gin.SetMode(viper.GetString("runmode"))

	// 创建 gin 实例
	g := gin.New()

	// 注册 404 路由
	g.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": 10003, "message": "Page not found."})
	})

	// 注册 /healthz handler
	g.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// 创建 HTTP Server 实例
	httpsrv := &http.Server{Addr: viper.GetString("addr"), Handler: g}

	log.Infow("start to listening the incoming requrests on http address", "addr", viper.GetString("addr"))

	if err := httpsrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Errorw(err.Error())
	}

	return nil

	return nil
}
