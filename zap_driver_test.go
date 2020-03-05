package slf4go_zap

import (
	"testing"

	slog "github.com/go-eden/slf4go"
	"go.uber.org/zap"
)

func TestInit(t *testing.T) {
	atomLevel := zap.NewAtomicLevelAt(zap.DebugLevel)
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	cfg := zap.Config{
		Level:       atomLevel,
		Development: false,
		// DisableCaller:     true,
		DisableStacktrace: true,
		Encoding:          "console",
		EncoderConfig:     encoderConfig,
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stdout"},
		InitialFields:     map[string]interface{}{"foo": "bar"},
	}

	Init(&cfg)

	slog.Debug("hello")
	slog.Info("what???")
	slog.Warnf("warnning: %v", "surrender")

	l := slog.GetLogger()
	l.BindFields(slog.Fields{
		"key": "value",
	})
	l.Errorf("error!!! %v", 100)

	l2 := slog.NewLogger("new-logger")
	l2.Info("l2 info log")
}
