package log

import "go.uber.org/zap/zapcore"

// Options 包含日志相关的配置
type Options struct {
	// 是否打印文件名和代码行数
	DisableCaller bool
	// 是否禁止在 panic 及以上级别打印堆栈信息
	DisableStacktrace bool
	// 日志等级，可选值：debug，info，warn，error，dpanic，panic，fatal
	Level string
	// 显示格式，可选值：console，json
	Format string
	// 日志输出位置
	OutPutPaths []string
}

// NewOptions 用于创建一个带有默认参数的 options 对象
func NewOptions() *Options {
	return &Options{
		DisableCaller:     false,
		DisableStacktrace: false,
		Level:             zapcore.InfoLevel.String(),
		Format:            "console",
		OutPutPaths:       []string{"stdout"},
	}
}
