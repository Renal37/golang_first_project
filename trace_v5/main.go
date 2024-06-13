package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime/trace"
	"syscall"
	"time"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	n := make(chan int)
	go func() {
		n <- Start(ctx)
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-stop:
		cancel()
	case val := <-n:
		fmt.Printf("N is : %d", val)
		return
	}
	fmt.Printf("N is : %d", <-n)

}
