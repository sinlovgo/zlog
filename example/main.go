package main

import (
	"fmt"
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
	zlog.Infof("info num %d", 10)
	zlog.Debug("Debug can not format %v", "debug-info")
}
