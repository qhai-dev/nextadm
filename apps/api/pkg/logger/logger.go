package logger

import (
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger Logger
	once   sync.Once
	//verbosity uint32
)

type Logger struct {
	log               logr.Logger
	SetVerbosityLevel func(uint32)
	GetVerbosityLevel func() uint32
	Flush             func()
}

type runtime struct {
	v uint32
}

func (r *runtime) ZapV() zapcore.Level {
	// zapcore.Level
	return -zapcore.Level(atomic.LoadUint32(&r.v))
}

func (r *runtime) Enabled(level zapcore.Level) bool {
	return level >= r.ZapV()
}

func (r *runtime) SetVerbosityLevel(v uint32) {
	atomic.StoreUint32(&r.v, v)
}

func (r *runtime) GetVerbosityLevel() uint32 {
	return atomic.LoadUint32(&r.v)
}

func InitLog() {
	once.Do(func() {
		logger = NewLogger(5)
	})
}

func FlushLog() {
	if logger.log.IsZero() {
		return
	}
	logger.Flush()
}

func NewLogger(v uint32) Logger {
	r := &runtime{
		v,
	}

	encoderConfig := &zapcore.EncoderConfig{
		NameKey:    "logger",
		MessageKey: "message",
		TimeKey:    "timestamp",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		LevelKey: "level",
		EncodeLevel: func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
			// 只设计 ERROR 和 INFO 两种级别日志
			// INFO 级别日志额外单独用v(n)去区分级别
			if level == zapcore.ErrorLevel {
				enc.AppendString(level.CapitalString())
			} else {
				enc.AppendString("INFO")
			}
		},
		CallerKey:     "caller",
		EncodeCaller:  zapcore.ShortCallerEncoder,
		FunctionKey:   "func",
		StacktraceKey: "stack",
	}

	encoder := zapcore.NewJSONEncoder(*encoderConfig)

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel && r.Enabled(lvl)
	})

	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel && r.Enabled(lvl)
	})

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.Lock(zapcore.AddSync(os.Stderr)), highPriority),
		zapcore.NewCore(encoder, zapcore.Lock(zapcore.AddSync(os.Stdout)), lowPriority),
	)

	zapLog := zap.New(
		core,
		zap.WithCaller(true),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	return Logger{
		log:               zapr.NewLoggerWithOptions(zapLog, zapr.LogInfoLevel("v"), zapr.ErrorKey("error")),
		SetVerbosityLevel: r.SetVerbosityLevel,
		GetVerbosityLevel: r.GetVerbosityLevel,
		Flush: func() {
			_ = zapLog.Sync()
		},
	}
}

func Info(msg string, keysAndValues ...any) {
	logger.log.Info(msg, keysAndValues...)
}

func Error(err error, msg string, keysAndValues ...any) {
	logger.log.Error(err, msg, keysAndValues...)
}

func V(v int) logr.Logger {
	return logger.log.V(v)
}

func SetVerbosityLevel(v uint32) {
	logger.SetVerbosityLevel(v)
}

func GetVerbosityLevel() uint32 {
	return logger.GetVerbosityLevel()
}

func WithName(name string) {
	logger.log = logger.log.WithName(name)
}
