package test

import (
	"os"
	"os/signal"
	"syscall"
)

var (
	onlyOneSigalHandler = make(chan struct{})
	shutdownSignals     = []os.Signal{os.Interrupt, syscall.SIGTERM}
)

func SetSignalHandler() (stopChan <-chan struct{}) {
	close(onlyOneSigalHandler) //panic when called twice

	stop := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		close(stop)
		<-c
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}
