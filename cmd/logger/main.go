package main

import (
	"context"
	"sync"
	"time"

	"github.com/wangfeiping/log"
)

func main() {

	defer log.Flush()

	log.Trace("init...")
	cancel := doLog()
	log.Info("ok, done.")

	keepRunning(cancel)
}

func output() {
	t := time.Now().Format(time.RFC3339Nano)
	log.Debugf("debug: %s", t)
	log.Warn("warn...")
	log.Error("error...")
}

func doLog() (cancel context.CancelFunc) {
	t := time.NewTicker(time.Duration(1) * time.Second)
	running := true
	var wg sync.WaitGroup
	cancel = func() {
		running = false
		t.Stop()
		wg.Wait()
	}
	go func() {
		wg.Add(1)
		output()
		for running {
			select {
			case <-t.C:
				{
					output()
				}
			default:
				{
					time.Sleep(100 * time.Millisecond)
				}
			}
		}
		wg.Done()
	}()
	return
}
