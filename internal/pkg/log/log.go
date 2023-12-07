package log

import (
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 确保 zapLogger 实现了 Logger
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

	// 如果传入的日记级别非法，则使用info级别
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}

	// 创建一个默认的encoder配置
	encoderConfig := zap.NewProductionEncoderConfig()
	// 自定义 MessageKey
	encoderConfig.MessageKey = "message"
	// 自定义 TimeKey 为 timestamp
	encoderConfig.TimeKey = "timestamp"
	// 指定时间序列化格式，将时间序列化为 `2006-01-02 15:04:05.000`
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	// 指定 time.Duration 序列化函数，将 time.Duration 序列化为经过的毫秒数的浮点数
	encoderConfig.EncodeDuration = func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendFloat64(float64(d) / float64(time.Millisecond))
	}

	// 创建 zap.Logger 需要的配置
	cfg := &zap.Config{
		DisableCaller:     opts.DisableCaller,
		DisableStacktrace: opts.DisableStacktrace,
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Encoding:          opts.Format,
		EncoderConfig:     encoderConfig,
		OutputPaths:       opts.OutPutPaths,
		ErrorOutputPaths:  []string{"stderr"},
	}

	// 使用 cfg 创建 *zap.Logger 对象
	z, err := cfg.Build(zap.AddStacktrace(zapcore.PanicLevel), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	logger := &zapLogger{z: z}

	zap.RedirectStdLog(z)

	return logger
}

func (l *zapLogger) Infow(msg string, keysAndValues ...any) {
	l.z.Sugar().Infow(msg, keysAndValues)
}
func Infow(msg string, keysAndValues ...any) {
	std.z.Sugar().Infow(msg, keysAndValues)
}

func (l *zapLogger) Warnw(msg string, keysAndValues ...any) {
	l.z.Sugar().Warnw(msg, keysAndValues)
}

func Warnw(msg string, keysAndValues ...any) {
	std.z.Sugar().Warnw(msg, keysAndValues)
}

func (l *zapLogger) Errorw(msg string, keysAndValues ...any) {
	l.z.Sugar().Errorw(msg, keysAndValues)
}

func Errorw(msg string, keysAndValues ...any) {
	std.z.Sugar().Errorw(msg, keysAndValues)
}

func (l *zapLogger) Panicw(msg string, keysAndValues ...any) {
	l.z.Sugar().Panicw(msg, keysAndValues)
}

func Panicw(msg string, keysAndValues ...any) {
	std.z.Sugar().Panicw(msg, keysAndValues)
}

func (l *zapLogger) Fatalw(msg string, keysAndValues ...any) {
	l.z.Sugar().Fatalw(msg, keysAndValues)
}

func Fatalw(msg string, keysAndValues ...any) {
	std.z.Sugar().Fatalw(msg, keysAndValues)
}

func (l *zapLogger) Debugw(msg string, keysAndValues ...any) {
	l.z.Sugar().Debugw(msg, keysAndValues)
}

func Debugw(msg string, keysAndValues ...any) {
	std.z.Sugar().Debugw(msg, keysAndValues)
}

// Sync 调用底层 zap.Logger 的 Sync 方法，将缓存中的日志刷新到磁盘文件中
func Sync() { std.Sync() }

func (l *zapLogger) Sync() {
	_ = l.z.Sync()
}
