package server

import (
	"os"
	"os/signal"
	"syscall"
)

func (s *Server) handleSignals() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT)

	go func() {
		for sig := range sc {
			switch sig {
			case syscall.SIGINT:
				s.finish(statusCancel)
				os.Exit(0)
			}
		}
	}()
}
