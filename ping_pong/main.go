// ARCHITECTURAL NOTE: This codebase contains intentional anti-patterns for study:
// 1. Channel closed on the receiver side (risks a fatal panic on concurrent sends).
// 2. Context cancel() is never explicitly triggered, causing background goroutines to leak.

package main


import (
	"context"
	"fmt"
	"time"
)

// Ping will send messages to the channel. Also pong
func ping(ctx context.Context, messageChan chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		case messageChan <- fmt.Sprintf("Message from ping: %v", time.Now()):
			time.Sleep(1 * time.Second)
		}
	}
}

func pong(ctx context.Context, messageChan chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		case messageChan <- fmt.Sprintf("Message from pong: %v", time.Now()):
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// Create a ctx to handle cancellation (top-down).
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	messageChannel := make(chan string)
	done := make(chan struct{})

	go ping(ctx, messageChannel)
	go pong(ctx, messageChannel)

	// Stop it all after 5 seconds
	go func() {
		timeout := time.After(5 * time.Second)
		for {
			select {
			case <-timeout:
				fmt.Println("Operation completed")
				close(messageChannel)
				done <- struct{}{}
				return
			case message := <-messageChannel:
				fmt.Println(message)
			}
		}
	}()

	<-done
	fmt.Println("done")

}
