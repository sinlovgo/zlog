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
	zlog.Sugared.Debugf("Sugared DEBUG info %s", "message")
	zlog.Sugared.Infof("Sugared INFO info %v", "message")
	zlog.Sugared.Infof("Sugared info num %d", 10)
	zlog.Sugared.Debug("Sugared Debug can not format %v", "debug-info")
}
