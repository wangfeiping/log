package main

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/wangfeiping/log"
)

func main() {
	defer log.Flush()

	cobra.EnableCommandSorting = false

	rootCmd := &cobra.Command{
		Use:               "logger",
		Short:             "Just for logger testing",
		PersistentPreRunE: initConfig,
		RunE:              doLogs,
	}
	rootCmd.PersistentFlags().String(log.FlagLogFile, "./test.log", "log file path")
	rootCmd.PersistentFlags().Int(log.FlagLogSize, 10, "log size(MB)")
	rootCmd.PersistentFlags().Int(log.FlagLogBackup, 10, "log backup")

	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(-1)
	}
}

func initConfig(cmd *cobra.Command, _ []string) error {
	viper.BindPFlags(cmd.Flags())
	log.Config(nil)
	log.Infof("starting at %s", getExecPath())
	return nil
}

func doLogs(_ *cobra.Command, _ []string) error {
	log.Trace("init...")
	cancel := doLog()
	log.Info("ok, done.")

	keepRunning(cancel)
	return nil
}

// getExecPath returns the execution path
func getExecPath() (execPath string) {
	file, _ := exec.LookPath(os.Args[0])
	execFile := filepath.Base(file)
	execPath, _ = filepath.Abs(file)
	if len(execPath) > 1 {
		rs := []rune(execPath)
		execPath = string(rs[0:(len(execPath) - len(execFile))])
	}
	return
}

func output() {
	t := time.Now().Format(time.RFC3339Nano)
	log.Debugf("debug: %s", t)
	log.Debugz("debug message",
		zap.Int64("timestamp", time.Now().UnixNano()),
		zap.String("for", "test"))
	log.Debug("debug: ", "time: ", time.Now().UnixNano())
	log.Warn("warn...")
	log.Warn("warn: ", "name: ", "test", "time: ", time.Now().UnixNano())
	log.Error("error...")
}

func doLog() (cancel context.CancelFunc) {
	t := time.NewTicker(time.Duration(100) * time.Millisecond)
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
					time.Sleep(10 * time.Millisecond)
				}
			}
		}
		wg.Done()
	}()
	return
}
