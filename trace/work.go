package main

import (
	"context"
	"time"
	"fmt"
)


func Start(ctx context.Context) {
	var n int
	ticker := time.Tick(100 * time.Millisecond)
	for {
		select {
        case <-ticker:
            n++
        case <-ctx.Done():
            fmt.Println(n)
            return
        }
	}
}
