package log

import (
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	mu sync.Mutex

	std = NewLogger(NewOptions())
)

var _ Logger = &zapLogger{}

type Logger interface {
	Debugw(msg string, keysAndValues ...any)
	Infow(msg string, keysAndValues ...any)
	Warnw(msg string, keysAndValues ...any)
	Errorw(msg string, keysAndValues ...any)
	Panicw(msg string, keysAndValues ...any)
	Fatalw(msg string, keysAndValues ...any)
	Sync()
}

func Init(opts *Options) {
	mu.Lock()
	defer mu.Unlock()

	std = NewLogger(opts)
}

type zapLogger struct {
	z *zap.Logger
}

func NewLogger(opts *Options) *zapLogger {

	if opts == nil {
		opts = NewOptions()
	}

	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}

	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.TimeKey = "timestamp"
	encodeConfig.MessageKey = "message"
	encodeConfig.EncodeTime = func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
		pae.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	encodeConfig.EncodeDuration = func(d time.Duration, pae zapcore.PrimitiveArrayEncoder) {
		pae.AppendFloat64(float64(d) / float64(time.Millisecond))
	}

	cfg := zap.Config{
		DisableCaller:     opts.DisableCaller,
		DisableStacktrace: opts.DisableStacktrace,
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Encoding:          opts.Format,
		EncoderConfig:     encodeConfig,
		OutputPaths:       opts.OutputPaths,
		ErrorOutputPaths:  []string{"stderr", "./test.log"},
	}

	zap, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return &zapLogger{
		z: zap,
	}
}

// Debugw implements Logger.
func (l *zapLogger) Debugw(msg string, keysAndValues ...any) {
	l.z.Sugar().Debugw(msg, keysAndValues...)
}

func Debugw(msg string, keysAndValues ...any) {
	std.z.Sugar().Debugw(msg, keysAndValues...)
}

// Errorw implements Logger.
func (l *zapLogger) Errorw(msg string, keysAndValues ...any) {
	l.z.Sugar().Errorw(msg, keysAndValues...)
}
func Errorw(msg string, keysAndValues ...any) {
	std.z.Sugar().Errorw(msg, keysAndValues...)
}

// Fatalw implements Logger.
func (l *zapLogger) Fatalw(msg string, keysAndValues ...any) {
	l.z.Sugar().Fatalw(msg, keysAndValues...)
}
func Fatalw(msg string, keysAndValues ...any) {
	std.z.Sugar().Fatalw(msg, keysAndValues...)
}

// Infow implements Logger.
func (l *zapLogger) Infow(msg string, keysAndValues ...any) {
	l.z.Sugar().Infow(msg, keysAndValues...)
}
func Infow(msg string, keysAndValues ...any) {
	std.z.Sugar().Infow(msg, keysAndValues...)
}

// Panicw implements Logger.
func (l *zapLogger) Panicw(msg string, keysAndValues ...any) {
	l.z.Sugar().Panicw(msg, keysAndValues...)
}
func Panicw(msg string, keysAndValues ...any) {
	std.z.Sugar().Panicw(msg, keysAndValues...)
}

// Warnw implements Logger.
func (l *zapLogger) Warnw(msg string, keysAndValues ...any) {
	l.z.Sugar().Warnw(msg, keysAndValues...)
}
func Warnw(msg string, keysAndValues ...any) {
	std.z.Sugar().Warnw(msg, keysAndValues...)
}

// Sync implements Logger.
func (l *zapLogger) Sync() {
	l.z.Sugar().Sync()
}
func Sync(msg string, keysAndValues ...any) {
	std.z.Sugar().Sync()
}
