package slf4go_zap

import (
	"testing"

	slog "github.com/go-eden/slf4go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestInit(t *testing.T) {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	atomLevel := zap.NewAtomicLevelAt(zap.DebugLevel)
	cfg := zap.Config{
		Level:             atomLevel,
		Development:       false,
		DisableCaller:     true,
		DisableStacktrace: true,
		Encoding:          "json",
		EncoderConfig:     encoderConfig,
		OutputPaths:       []string{"stdout", "/tmp/stdout"},
		ErrorOutputPaths:  []string{"stdout", "/tmp/zaperr"},
		InitialFields:     map[string]interface{}{"foo": "bar"},
	}

	Init(&cfg)

	slog.Debug("hello")
	slog.Info("what???")
	slog.Warnf("warnning: %v", "surrender")

	l := slog.GetLogger()
	l.BindFields(slog.Fields{
		"logger": "test",
	})
	l.Errorf("error!!! %v", 100)
}
