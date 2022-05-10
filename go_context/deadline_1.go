package main

import (
	"context"
	"fmt"
	_ "net/http/pprof"
	"time"
)

// 使用 WithTimeout ctx 控制协程的运行时间
func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(800*time.Millisecond))
	defer cancel()

	go handle2(ctx, 500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle2(ctx context.Context, duration time.Duration) {
	// 如果协程先退出，则不会执行携程内的“handle”
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}
