package log

import (
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ Logger = &zapLogger{}

var (
	mu  sync.Mutex
	std = NewLogger(NewOptions())
)

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

	std = NewLogger(opts)
}

func NewLogger(opts *Options) *zapLogger {
	if opts == nil {
		opts = NewOptions()
	}

	// 如果是非法 level，默认使用info
	var zapLever zapcore.Level
	if err := zapLever.UnmarshalText([]byte(opts.Level)); err != nil {
		zapLever = zapcore.InfoLevel
	}

	// 处理输出格式
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.MessageKey = "message"
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
		pae.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	encoderConfig.EncodeDuration = func(d time.Duration, pae zapcore.PrimitiveArrayEncoder) {
		pae.AppendFloat64(float64(d) / float64(time.Millisecond))
	}

	cfg := &zap.Config{
		DisableCaller:     opts.DisableCaller,
		DisableStacktrace: opts.DisableStacktrace,
		Level:             zap.NewAtomicLevelAt(zapLever),
		Encoding:          opts.Format,
		EncoderConfig:     encoderConfig,
		OutputPaths:       opts.OutputPaths,
		ErrorOutputPaths:  []string{"stderr"},
	}
	z, err := cfg.Build(zap.AddStacktrace(zapcore.DPanicLevel), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	logger := &zapLogger{z: z}

	return logger
}

func (l *zapLogger) Debugw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Debugw(msg, keysAndValues)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Debugw(msg, keysAndValues)
}

func (l *zapLogger) Infow(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Infow(msg, keysAndValues)
}

func Infow(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Infow(msg, keysAndValues)
}

func (l *zapLogger) Warnw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Warnw(msg, keysAndValues)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Warnw(msg, keysAndValues)
}

func (l *zapLogger) Errorw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Error(msg, keysAndValues)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Error(msg, keysAndValues)
}

func (l *zapLogger) Panicw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Panicw(msg, keysAndValues)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Panicw(msg, keysAndValues)
}

func (l *zapLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Fatalw(msg, keysAndValues)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Fatalw(msg, keysAndValues)
}

func (l *zapLogger) Sync() {
	_ = l.z.Sync()
}

func Sync() {
	_ = std.z.Sync()
}
