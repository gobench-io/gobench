package gobench

import (
	"os"
	"os/signal"
	"syscall"
)

func (c *Collect) handleSignals() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT)

	go func() {
		for sig := range sc {
			switch sig {
			case syscall.SIGINT:
				c.finish(statusCancel)
				os.Exit(0)
			}
		}
	}()
}
