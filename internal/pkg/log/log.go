package log

import (
	"context"
	"miniblog/internal/pkg/known"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debugw(msg string, keysAndValues ...any)
	Infow(msg string, keysAndValues ...any)
	Warnw(msg string, keysAndValues ...any)
	Errorw(msg string, keysAndValues ...any)
	Panicw(msg string, keysAndValues ...any)
	Fatalw(msg string, keysAndValues ...any)
	Sync()
}

var _ Logger = &zapLogger{}

type zapLogger struct {
	z *zap.Logger
}

var (
	mu  sync.Mutex
	std = NewLogger(NewOptions())
)

// 初始化全局 logger
func Init(opts *Options) {

	mu.Lock()
	defer mu.Unlock()

	std = NewLogger(opts)
}

// 通过传入的 options 构建 zapLogger
func NewLogger(opts *Options) *zapLogger {

	if opts == nil {
		opts = NewOptions()
	}

	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}

	encoder := zap.NewProductionEncoderConfig()
	encoder.MessageKey = "message"
	encoder.TimeKey = "timestamp"
	encoder.EncodeTime = func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
		pae.AppendString(t.Format(time.DateTime))

	}

	cfg := &zap.Config{
		DisableCaller:     opts.DisableCaller,
		DisableStacktrace: opts.DisableStacktrace,
		Level:             zap.NewAtomicLevelAt(zapLevel),
		OutputPaths:       opts.OutputPaths,
		Encoding:          opts.Format,
		EncoderConfig:     encoder,
		ErrorOutputPaths:  []string{"miniblog_err.log"},
	}
	z, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	// 把标准库的 log.Logger 的 info 级别的输出重定向到 zap.Logger
	zap.RedirectStdLog(z)

	return &zapLogger{z: z}
}

func C(ctx context.Context) *zapLogger {
	return std.C(ctx)
}

func (l *zapLogger) C(ctx context.Context) *zapLogger {
	lc := l.clone()

	if requestID := ctx.Value(known.XRequestIDKey); requestID != nil {
		lc.z = lc.z.With(zap.Any(known.XRequestIDKey, requestID))
	}

	return lc
}

func (l *zapLogger) clone() *zapLogger {
	lc := *l
	return &lc
}

// Debugw
func (l *zapLogger) Debugw(msg string, keysAndValues ...any) {
	l.z.Sugar().Debugw(msg, keysAndValues...)
}

func Debugw(msg string, keysAndValues ...any) {
	std.z.Sugar().Debugw(msg, keysAndValues...)
}

// Errorw
func (l *zapLogger) Errorw(msg string, keysAndValues ...any) {
	l.z.Sugar().Errorw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...any) {
	std.z.Sugar().Errorw(msg, keysAndValues...)
}

// Fatalw
func (l *zapLogger) Fatalw(msg string, keysAndValues ...any) {
	l.z.Sugar().Fatalw(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...any) {
	std.z.Sugar().Fatalw(msg, keysAndValues...)
}

// Infow
func (l *zapLogger) Infow(msg string, keysAndValues ...any) {
	l.z.Sugar().Infow(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...any) {
	std.z.Sugar().Infow(msg, keysAndValues...)
}

// Panicw
func (l *zapLogger) Panicw(msg string, keysAndValues ...any) {
	l.z.Sugar().Panicw(msg, keysAndValues...)
}
func Panicw(msg string, keysAndValues ...any) {
	std.z.Sugar().Panicw(msg, keysAndValues...)
}

// Warnw
func (l *zapLogger) Warnw(msg string, keysAndValues ...any) {
	l.z.Sugar().Warnw(msg, keysAndValues...)
}
func Warnw(msg string, keysAndValues ...any) {
	std.z.Sugar().Warnw(msg, keysAndValues...)
}

// Sync
func (l *zapLogger) Sync() {
	_ = l.z.Sync()
}

func Sync() {
	std.Sync()
}
