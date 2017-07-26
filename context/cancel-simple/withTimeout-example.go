package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)

	defer cancel() // needs to be called

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
