package log

import (
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	std = NewZapLogger(NewOptions())
	mu  sync.Mutex
)

var _ Logger = &zapLogger{}

type Logger interface {
	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	Sync()
}

type zapLogger struct {
	z *zap.Logger
}

func Init(opts *Options) {
	mu.Lock()
	defer mu.Unlock()

	std = NewZapLogger(opts)
}

func NewZapLogger(opts *Options) *zapLogger {
	if opts == nil {
		opts = NewOptions()
	}

	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.MessageKey = "message"
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
		pae.AppendString(t.Format(time.DateTime))
	}

	cfg := &zap.Config{
		DisableCaller:     opts.DisableCaller,
		DisableStacktrace: opts.DisableStacktrace,
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Encoding:          opts.Format,
		EncoderConfig:     encoderConfig,
		OutputPaths:       opts.OutputPaths,
		ErrorOutputPaths:  []string{"stderr"},
	}

	z, err := cfg.Build(zap.AddStacktrace(zapcore.PanicLevel), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	logger := &zapLogger{z: z}

	zap.RedirectStdLog(z)

	return logger
}

// Sync 调用底层 zap.Logger 的 Sync 方法，将缓存中的日志刷新到磁盘文件中. 主程序需要在退出前调用 Sync.
func Sync() { std.Sync() }

func (l *zapLogger) Sync() {
	_ = l.z.Sync()
}

// Debugw 输出 debug 级别的日志.
func Debugw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Debugw(msg, keysAndValues...)
}

func (l *zapLogger) Debugw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Debugw(msg, keysAndValues...)
}

// Infow 输出 info 级别的日志.
func Infow(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Infow(msg, keysAndValues...)
}

func (l *zapLogger) Infow(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Infow(msg, keysAndValues...)
}

// Warnw 输出 warning 级别的日志.
func Warnw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Warnw(msg, keysAndValues...)
}

func (l *zapLogger) Warnw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Warnw(msg, keysAndValues...)
}

// Errorw 输出 error 级别的日志.
func Errorw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Errorw(msg, keysAndValues...)
}

func (l *zapLogger) Errorw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Errorw(msg, keysAndValues...)
}

// Panicw 输出 panic 级别的日志.
func Panicw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Panicw(msg, keysAndValues...)
}

func (l *zapLogger) Panicw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Panicw(msg, keysAndValues...)
}

// Fatalw 输出 fatal 级别的日志.
func Fatalw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Fatalw(msg, keysAndValues...)
}

func (l *zapLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Fatalw(msg, keysAndValues...)
}
