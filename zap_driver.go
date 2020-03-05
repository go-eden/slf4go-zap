package slf4go_zap

import (
	slog "github.com/go-eden/slf4go"
	"go.uber.org/zap"
)

type ZapDriver struct {
	loggers map[string]*zap.Logger
	cfg     *zap.Config
}

func Init(cfg *zap.Config) {
	d := ZapDriver{}
	d.loggers = make(map[string]*zap.Logger, 0)

	d.cfg = cfg
	if logger, err := cfg.Build(); err != nil {
		panic(err)
	} else {
		d.loggers["global"] = logger
	}
	slog.SetDriver(&d)
}

func (d *ZapDriver) Name() string {
	return "slf4go-zap"
}

func (d *ZapDriver) Print(l *slog.Log) {
	pLogger := d.loggers["global"]
	// 处理field
	if l.Fields != nil {
		if _, ok := d.loggers[l.Logger]; !ok {
			fields := make([]zap.Field, 0)
			for k, v := range l.Fields {
				fields = append(fields, zap.Any(k, v))
			}
			d.loggers[l.Logger] = d.loggers["global"].With(fields...)
		}

		pLogger = d.loggers[l.Logger]
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
	case slog.FataLevel:
		if l.Format == nil {
			pLogger.Sugar().Fatal(l.Args...)
		} else {
			pLogger.Sugar().Fatalf(*l.Format, l.Args...)
		}
	}
}

func (d *ZapDriver) GetLevel(logger string) (sl slog.Level) {
	l := d.cfg.Level.Level()

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
		sl = slog.FataLevel
	default:
		sl = slog.TraceLevel
	}
	return
}
