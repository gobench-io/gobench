package master

import (
	"os"
	"os/signal"
	"syscall"
)

func (m *master) handleSignals() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT)

	go func() {
		for sig := range sc {
			switch sig {
			case syscall.SIGINT:
				m.finish(statusCancel)
				os.Exit(0)
			}
		}
	}()
}
