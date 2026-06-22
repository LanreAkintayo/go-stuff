package main

import (
	"fmt"
	"time"
)

/*
DRILL 1: THE BASIC HANDSHAKE (GOROUTINES + CHANNELS)
Question: Pass a single string message ("Hello from Ibadan!") from a background
worker goroutine to the main thread using an unbuffered channel and print it.
*/
func runDrillOne() {
	stringChannel := make(chan string)

	go func() {
		stringChannel <- "Hello from Ibadan"
	}()

	message := <-stringChannel
	fmt.Println(message)
}

/*
DRILL 2: THE LOOP & CLOSE DRILL (PREVENTING DEADLOCKS)
Question: Spawn a background goroutine that uses a loop to send numbers 10, 20, 30, 40, 50
into an integer channel. Close the channel properly to prevent a deadlock, and read
the data in the main thread using a "for range" loop.
*/
func runDrillTwo() {
	intChannel := make(chan int)

	go func() {
		for i := 10; i <= 50; i += 10 {
			intChannel <- i
		}
		close(intChannel)
	}()

	for values := range intChannel {
		fmt.Println("Values: ", values)
	}
}

/*
DRILL 3: THE HEALTH-CHECK PING DAEMON (TIMERS & LIFESPAN SHUTDOWN)
Question: Build an asynchronous monitoring system. It must log the current time every
1 second using a ticker, spawn a background worker that takes 500ms to process a response,
and use a central select block to completely shut down the application after exactly 10 seconds.
*/
func runDrillThree() {
	shutdownChan := time.After(10 * time.Second)
	tickerChan := time.Tick(1 * time.Second)

	fmt.Println("Starting monitoring session")

	for {
		select {
		case currentTime := <-tickerChan:
			fmt.Println("Current time: ", currentTime.Format("15:04:05"))

			go func() {
				time.Sleep(500 * time.Millisecond)
				fmt.Printf("Response for %v received successfully!\n", currentTime.Format("15:04:05"))
			}()

		case shutdownTime := <-shutdownChan:
			fmt.Println("Monitoring session ended at exactly: ", shutdownTime.Format("15:04:05"))
			return
		}
	}
}


/*
DRILL 4: This is me testing the behaviour of channels on the main thread
*/
func runDrillFour() {
	// The goal is not to make the main thread execute until I pass some data to a channel

	testChan := make(chan int)

	// go func(){
	// 	testChan <- 4
	// }()

	fmt.Println("About to listen to chan")
	<-testChan
	fmt.Println("This is the end")
}
