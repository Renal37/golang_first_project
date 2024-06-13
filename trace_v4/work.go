package main

import (
	"context"
	"time"
)

func Start(ctx context.Context) int {
	var n int
	ticker := time.Tick(100 * time.Millisecond)
	for {
		select {
		case <-ticker:
			n++
		case <-ctx.Done():
			return n
		}
	}
}
