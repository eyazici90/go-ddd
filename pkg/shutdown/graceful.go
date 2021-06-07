package shutdown

import (
	"os"
	"os/signal"
)

func Gracefully() {
	quit := make(chan os.Signal, 1)
	defer close(quit)

	signal.Notify(quit, os.Interrupt)
	<-quit
}
