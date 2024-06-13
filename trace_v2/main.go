package main

import (
    "context"
    "fmt"
    "os"
    "os/signal"
    "sync"
    "syscall"
)


func main() {
    ctx, cancel := context.WithCancel(context.Background())
    var n int
    mx := new(sync.Mutex)
    go func() {
        mx.Lock()
        n = Start(ctx)
        mx.Unlock()
    }()
    stop := make(chan os.Signal, 1)
    signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

    <-stop
    cancel()
    mx.Lock()
    fmt.Printf("N is : %d", n)
    mx.Unlock()
}