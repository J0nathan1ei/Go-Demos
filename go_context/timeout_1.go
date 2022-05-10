package main

import (
	"context"
	"fmt"
	_ "net/http/pprof"
	"time"
)

// 使用 WithTimeout ctx 控制协程的运行时间
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 800*time.Millisecond)
	defer cancel()

	go handle(ctx, 500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	// 如果协程先退出，则不会执行携程内的“handle”
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}
