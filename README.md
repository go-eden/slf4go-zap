# slf4go-zap
This is a default Driver implementation for bridging `slf4go` and `zap`.

# Install #

`slf4go-zap` dependents on `slf4go` and `zap`.

``` shell
go get github.com/go-eden/slf4go-zap
```

# Usage

`slf4go-zap` focuses on bridging logs, you should configure `zap` according to your needs.

zap_driver_test.go This is a simple example of `slf4go-zap`:


# Notice

Only support zap.SugaredLogger, so this library don't have lots of features currently.

Hope you can help me improve this library, any `Pull Request` will be very welcomed.
