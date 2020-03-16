[![TravisBuildStatus](https://api.travis-ci.org/sinlovgo/zlog.svg?branch=master)](https://travis-ci.org/sinlovgo/zlog)
[![GoDoc](https://godoc.org/github.com/sinlovgo/zlog?status.png)](https://godoc.org/github.com/sinlovgo/zlog/)
[![GoReportCard](https://goreportcard.com/badge/github.com/sinlovgo/zlog)](https://goreportcard.com/report/github.com/sinlovgo/zlog)
[![codecov](https://codecov.io/gh/sinlovgo/zlog/branch/master/graph/badge.svg)](https://codecov.io/gh/sinlovgo/zlog)

## for what

- this project use viper read config file for log
- config reader use [github.com/spf13/viper](https://github.com/spf13/viper)
- log use [go.uber.org/zap](https://github.com/uber-go/zap)
- rotation use [gopkg.in/natefinch/lumberjack.v2](https://github.com/natefinch/lumberjack)

## depends

in go mod project

```bash
# before above global settings
# test version info
$ git ls-remote -q http://github.com/sinlovgo/zlog.git

# test depends see full version
$ go list -v -m -versions github.com/sinlovgo/zlog
# or use last version add go.mod by script
$ echo "go mod edit -require=$(go list -m -versions github.com/sinlovgo/zlog.git | awk '{print $1 "@" $NF}')"
$ echo "go mod vendor"
```

### config file

```yaml
zap:
  AtomicLevel: -1 # DebugLevel:-1 InfoLevel:0 WarnLevel:1 ErrorLevel:2
  FieldsAuto: false # is use auto Fields key set
  Fields:
    Key: key
    Val: val
  Development: true # is open Open file and line number
  Encoding: console # output format, only use console or json, default is console
  rotate:
    Filename: log/zlog.log # Log file path
    MaxSize: 16 # Maximum size of each log file, Unit: M
    MaxBackups: 10 # How many backups are saved in the log file
    MaxAge: 7 # How many days can the file be keep, Unit: day
    Compress: true # need compress
  EncoderConfig:
    TimeKey: time
    LevelKey: level
    NameKey: logger
    CallerKey: caller
    MessageKey: msg
    StacktraceKey: stacktrace
```

### use-code

```go
import (
	"github.com/sinlovgo/zlog"
)
func main() {
    // default config conf/config.yaml
	err := zlog.Init("")
	if err != nil {
		fmt.Printf("zlog init error %v\n", err)
		return
	}
	zlog.Debugf("info %s", "message")
	zlog.Infof("info %v", "message")
}
```

## dev

- use to replace
 `sinlovgo/zlog` to you code

- and run

```bash
make init
```

- test code

```bash
make test
```

add main.go file and run

```bash
make run
```

# dev

## evn

- golang sdk 1.13+

## docker

base docker file can replace
