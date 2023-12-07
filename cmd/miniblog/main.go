// Copyright 2022 Innkeeper Belm(孔令飞) <nosbelm@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/marmotedu/miniblog.

package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	url := "http://baidu.com"
	logger.Info("failed to fetch url",
		zap.String("url", url),
		zap.Int("accept", 3),
		zap.Duration("backoff", time.Second),
	)

	sugar := logger.Sugar()
	sugar.Infow("failed to fetch url", "url", url, "accept", 3)

	// command := miniblog.NewMiniBlogCommand()
	// if err := command.Execute(); err != nil {
	// 	os.Exit(1)
	// }
}
