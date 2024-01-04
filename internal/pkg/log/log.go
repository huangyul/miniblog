package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
	"time"
)

// 确保 zapLogger 实现了 Logger 接口
var _ Logger = &zapLogger{}

// Logger 定义 miniblog 项目的日志接口
type Logger interface {
	Debugw(msg string, keysAndValues ...any)
	Infow(msg string, keysAndValues ...any)
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	Sync()
}

// zapLogger 是 Logger
type zapLogger struct {
	z *zap.Logger
}

func (l *zapLogger) Debugw(msg string, keysAndValues ...any) {
	l.z.Sugar().Debugw(msg, keysAndValues)
}

func Debugw(msg string, keysAndValues ...any) {
	std.z.Sugar().Debugw(msg, keysAndValues)
}

func (l *zapLogger) Infow(msg string, keysAndValues ...any) {
	l.z.Sugar().Infow(msg, keysAndValues)
}

func Infow(msg string, keysAndValues ...any) {
	std.z.Sugar().Infow(msg, keysAndValues)
}

func (l *zapLogger) Warnw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Warnw(msg, keysAndValues)
}

func Warnw(msg string, keysAndValues ...any) {
	std.z.Sugar().Warnw(msg, keysAndValues)
}

func (l *zapLogger) Errorw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Errorw(msg, keysAndValues)
}

func Errorw(msg string, keysAndValues ...any) {
	std.z.Sugar().Errorw(msg, keysAndValues)
}

func (l *zapLogger) Panicw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Panicw(msg, keysAndValues)
}

func Panicw(msg string, keysAndValues ...any) {
	std.z.Sugar().Panicw(msg, keysAndValues)
}

func (l *zapLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Fatalw(msg, keysAndValues)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Fatalw(msg, keysAndValues)

}

// Sync 将缓存中的日志刷新到磁盘中，主程序需要在退出前调用
func (l *zapLogger) Sync() {
	_ = l.z.Sync()
}

func Sync() { std.Sync() }

var (
	mu sync.Mutex
	// 定义了默认的全局 Logger
	std = NewLogger(NewOptions())
)

// Init 使用指定 Options 初始化全局 Logger
func Init(opts *Options) {
	mu.Lock()
	defer mu.Unlock()

	std = NewLogger(opts)
}

// NewLogger 根据传入的 Options 创建局部的 Logger
func NewLogger(opts *Options) *zapLogger {
	if opts == nil {
		opts = NewOptions()
	}

	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}

	// 创建一个默认的 encoder 配置
	encoderConfig := zap.NewProductionEncoderConfig()
	// 自定义 MessageKey 为 message
	encoderConfig.MessageKey = "message"
	// 自定义 TimeKey 为 timestamp
	encoderConfig.TimeKey = "timestamp"
	// 指定时间序列化函数
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("2006-01-02 15:04:05.000"))
	}
	// 指定 time.Duration 序列化函数
	encoderConfig.EncodeDuration = func(duration time.Duration, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendFloat64(float64(duration) / float64(time.Millisecond))
	}

	// 创建构建 zap.Logger 需要的配置
	cfg := &zap.Config{
		DisableCaller:     opts.DisableCaller,
		DisableStacktrace: opts.DisableStacktrace,
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Encoding:          opts.Format,
		EncoderConfig:     encoderConfig,
		OutputPaths:       opts.OutputPaths,
		ErrorOutputPaths:  []string{"stderr"},
	}

	// 使用 cfg 创建 *zap.Logger 对象
	z, err := cfg.Build(zap.AddStacktrace(zapcore.PanicLevel), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	logger := &zapLogger{z: z}

	// 把标准库的 log.Logger 的 info 级别输出重定向到 zap.Logger
	zap.RedirectStdLog(z)

	return logger
}
