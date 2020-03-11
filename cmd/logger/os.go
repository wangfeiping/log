package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/wangfeiping/log"
)

func keepRunning(cancel context.CancelFunc) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	select {
	case s, ok := <-signals:
		log.Debugf("system signal [%v] %t, trying to run cancel...", s, ok)
		if !ok {
			break
		}
		if cancel != nil {
			cancel()
		}
		log.Flush()
		os.Exit(1)
	}
}
