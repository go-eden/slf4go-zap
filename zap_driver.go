package slf4gozap

import (
	slog "github.com/go-eden/slf4go"
	"go.uber.org/zap"
)

// SkipUntilTrueCaller is the skip level which prints out the actual caller instead of slf4go or slf4go-zap wrappers
const SkipUntilTrueCaller = 3

// Config wraps zap config with some custom options
type Config struct {
	ZapConfig  *zap.Config
	ZapOptions []zap.Option
}

// ZapDriver is the wrapper around zap logger and its config
type ZapDriver struct {
	logger *zap.Logger
	cfg    *Config
}

// Init initializes the driver using the provided config wrapper
func Init(cfg *Config) {
	d := ZapDriver{}

	d.cfg = cfg
	var err error
	if d.logger, err = cfg.ZapConfig.Build(cfg.ZapOptions...); err != nil {
		panic(err)
	}
	slog.SetDriver(&d)
}

// Name returns the driver's name, which in this case, "slf4go-zap"
func (d *ZapDriver) Name() string {
	return "slf4go-zap"
}

// Print specifies how the driver will actually printout the log
func (d *ZapDriver) Print(l *slog.Log) {
	pLogger := d.logger
	// 处理field
	if l.Fields != nil {
		fields := make([]zap.Field, 0, len(l.Fields))
		for k, v := range l.Fields {
			fields = append(fields, zap.Any(k, v))
		}
		pLogger = d.logger.With(fields...)
	}

	defer pLogger.Sync()
	switch l.Level {
	case slog.TraceLevel:
		if l.Format == nil {
			pLogger.Sugar().Debug(l.Args...)
		} else {
			pLogger.Sugar().Debugf(*l.Format, l.Args...)
		}
	case slog.DebugLevel:
		if l.Format == nil {
			pLogger.Sugar().Debug(l.Args...)
		} else {
			pLogger.Sugar().Debugf(*l.Format, l.Args...)
		}
	case slog.InfoLevel:
		if l.Format == nil {
			pLogger.Sugar().Info(l.Args...)
		} else {
			pLogger.Sugar().Infof(*l.Format, l.Args...)
		}
	case slog.WarnLevel:
		if l.Format == nil {
			pLogger.Sugar().Warn(l.Args...)
		} else {
			pLogger.Sugar().Warnf(*l.Format, l.Args...)
		}
	case slog.ErrorLevel:
		if l.Format == nil {
			pLogger.Sugar().Error(l.Args...)
		} else {
			pLogger.Sugar().Errorf(*l.Format, l.Args...)
		}
	case slog.PanicLevel:
		if l.Format == nil {
			pLogger.Sugar().Panic(l.Args...)
		} else {
			pLogger.Sugar().Panicf(*l.Format, l.Args...)
		}
	case slog.FatalLevel:
		if l.Format == nil {
			pLogger.Sugar().Fatal(l.Args...)
		} else {
			pLogger.Sugar().Fatalf(*l.Format, l.Args...)
		}
	}
}

// GetLevel returns the current level of the logger
func (d *ZapDriver) GetLevel(logger string) (sl slog.Level) {
	l := d.cfg.ZapConfig.Level.Level()

	switch l {
	case zap.DebugLevel:
		sl = slog.DebugLevel
	case zap.InfoLevel:
		sl = slog.InfoLevel
	case zap.WarnLevel:
		sl = slog.WarnLevel
	case zap.ErrorLevel:
		sl = slog.ErrorLevel
	case zap.DPanicLevel:
		sl = slog.PanicLevel
	case zap.PanicLevel:
		sl = slog.PanicLevel
	case zap.FatalLevel:
		sl = slog.FatalLevel
	default:
		sl = slog.TraceLevel
	}
	return
}
