package main

import (
	"time"

	"github.com/wangfeiping/log"
)

func main() {

	defer log.Flush()

	log.Trace("init...")

	t := time.Now().Format(time.RFC3339Nano)
	log.Debugf("debug: %s", t)
	log.Warn("warn...")
	log.Error("error...")
	log.Info("ok, done.")

	keepRunning(nil)
}
