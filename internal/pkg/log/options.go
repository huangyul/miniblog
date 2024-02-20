package log

import "go.uber.org/zap/zapcore"

type Options struct {
	// 是否开启 caller，开启后会在日志中显示文件和行号
	DisableCaller bool
	// 是否在 panic 以及以上打印堆栈信息
	DisableStacktrace bool
	// 指定日志级别
	Level string
	// 指定输出格式
	Format string
	// 指定输出位置
	OutputPaths []string
}

// NewOptions 创建一个带有默认值的 options 对象
func NewOptions() *Options {
	return &Options{
		DisableCaller:     false,
		DisableStacktrace: false,
		Level:             zapcore.InfoLevel.String(),
		Format:            "console",
		OutputPaths:       []string{"miniblog.log"},
	}
}
