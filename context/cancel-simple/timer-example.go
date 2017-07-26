package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// in an other thread, do some blocking check
	go func() {
		time.Sleep(time.Second)
		cancel() // out
	}()

	sleepAndTalk(ctx, time.Second*5, "Hello")
}

func sleepAndTalk(ctx context.Context, duration time.Duration, msg string) {
	select {
	case <-time.After(duration):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Print(ctx.Err())
	}
}
