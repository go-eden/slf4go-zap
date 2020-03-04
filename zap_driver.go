package slf4go_zap

import (

	// slog "github.com/phenix3443/slf4go"
	slog "github.com/go-eden/slf4go"
	"go.uber.org/zap"
)

type ZapDriver struct {
	logger *zap.Logger
	cfg    *zap.Config
}

func Init(cfg *zap.Config) {
	d := ZapDriver{}
	d.cfg = cfg
	if logger, err := cfg.Build(); err != nil {
		panic(err)
	} else {
		d.logger = logger

	}
	slog.SetDriver(&d)
}

func (d *ZapDriver) Name() string {
	return "zap"
}

func (d *ZapDriver) Print(l *slog.Log) {
	defer d.logger.Sugar().Sync()

	if l.Fields != nil {
		d.cfg.InitialFields = l.Fields
		var err error
		if d.logger, err = d.cfg.Build(); err != nil {
			panic(err)
		}
		l.Fields = nil
	}

	switch l.Level {
	case slog.TraceLevel:
		if l.Format == nil {
			d.logger.Sugar().Debug(l.Args...)
		} else {
			d.logger.Sugar().Debugf(*l.Format, l.Args...)
		}
	case slog.DebugLevel:
		if l.Format == nil {
			d.logger.Sugar().Debug(l.Args...)
		} else {
			d.logger.Sugar().Debugf(*l.Format, l.Args...)
		}
	case slog.InfoLevel:
		if l.Format == nil {
			d.logger.Sugar().Info(l.Args...)
		} else {
			d.logger.Sugar().Infof(*l.Format, l.Args...)
		}
	case slog.WarnLevel:
		if l.Format == nil {
			d.logger.Sugar().Warn(l.Args...)
		} else {
			d.logger.Sugar().Warnf(*l.Format, l.Args...)
		}
	case slog.ErrorLevel:
		if l.Format == nil {
			d.logger.Sugar().Error(l.Args...)
		} else {
			d.logger.Sugar().Errorf(*l.Format, l.Args...)
		}
	case slog.PanicLevel:
		if l.Format == nil {
			d.logger.Sugar().Panic(l.Args...)
		} else {
			d.logger.Sugar().Panicf(*l.Format, l.Args...)
		}
	case slog.FataLevel:
		if l.Format == nil {
			d.logger.Sugar().Fatal(l.Args...)
		} else {
			d.logger.Sugar().Fatalf(*l.Format, l.Args...)
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
