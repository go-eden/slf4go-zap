# slf4go-zap

This is a default Driver implementation for bridging `slf4go` and `zap`.

## Install

`slf4go-zap` dependents on `slf4go` and `zap`.

``` shell
go get github.com/go-eden/slf4go-zap
```

## Usage

`slf4go-zap` focuses on bridging logs, you should configure `zap` according to your needs.

```go
package main

import (
 slog "github.com/go-eden/slf4go"
 slogzap "github.com/go-eden/slf4go-zap"
 "go.uber.org/zap"
)

func main() {

  zapcfg = zap.Config{
  Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
  Development: false,
  // DisableCaller:     true,
  DisableStacktrace: true,
  Encoding:          "console",
  EncoderConfig:     zap.NewDevelopmentEncoderConfig(),
  OutputPaths:       []string{"stdout"},
  ErrorOutputPaths:  []string{"stdout"},
  InitialFields:     map[string]interface{}{"foo": "bar"},
 }

 cfg = slogzap.Config{
  ZapConfig: &zapcfg,
  ZapOptions: []zap.Option{
   zap.AddCallerSkip(slogzap.SkipUntilTrueCaller), // 3
  },
 }

 slogzap.Init(&cfg)

 // use the global logger
 slog.Debug("zap")

 // or create a new one and use it
 l := slog.GetLogger()
 l.Errorf("default logger name=%s", l.Name())

}

```

Further examples can be seen in the zap_driver_test.go file.

## Notice

Only support zap.SugaredLogger, so this library don't have lots of features currently.

zap.Option is now supported.

Hope you can help me improve this library, any `Pull Request` will be very welcomed.

## Contributor

@phenix3443 @mikeychowy
