package shutdown

import (
	"os"
	"os/signal"
	"syscall"
)

func Gracefully() {
	quit := make(chan os.Signal, 1)
	defer close(quit)

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
}
