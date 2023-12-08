// Copyright 2022 Innkeeper Belm(孔令飞) <nosbelm@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/marmotedu/miniblog.

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	g := gin.Default()

	g.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  10003,
			"message": "page not found",
		})
	})

	g.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  0,
			"message": "server is ok",
		})
	})

	g.Run(":3333")

	// logger, _ := zap.NewProduction()
	// defer logger.Sync()

	// url := "http://baidu.com"
	// logger.Info("failed to fetch url",
	// 	zap.String("url", url),
	// 	zap.Int("accept", 3),
	// 	zap.Duration("backoff", time.Second),
	// )

	// sugar := logger.Sugar()
	// sugar.Infow("failed to fetch url", "url", url, "accept", 3)

	// command := miniblog.NewMiniBlogCommand()
	// if err := command.Execute(); err != nil {
	// 	os.Exit(1)
	// }
}
