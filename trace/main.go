package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go Start(ctx)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	cancel()
	t := time.NewTimer(10 * time.Second)
	defer t.Stop()

	select {
	case <-t.C:
		fmt.Println("timeout")
	case <-stop:
		fmt.Println("stop")
	}
}
