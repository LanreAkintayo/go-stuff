package main


import (
	"fmt"
	"time"

)

func Timer(){
	timerChan := time.Tick(1 * time.Second)

	counter := 0

	for range timerChan{
		counter++
		fmt.Printf("Current Count: %ds\n", counter)
	}
}