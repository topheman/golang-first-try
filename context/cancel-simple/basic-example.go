package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// in an other thread, do some blocking check
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		// blocks until entry in Stdin
		scanner.Scan()
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
