package log

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	z *zap.Logger
}

var (
	mu  sync.Mutex
	std = NewLogger(NewOptions()) // 默认是全局的logger
)

// Init 指定选择初始化log
func Init(opts *Options) {
	mu.Lock()
	defer mu.Unlock()

	std = NewLogger(opts)

}

// NewLogger 传入配置，新建一个logger
func NewLogger(opts *Options) *zapLogger {
	if opts == nil {
		opts = NewOptions()
	}

	// ru'g'y
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}

	return &zapLogger{
		z: &zap.Logger{},
	}
}
