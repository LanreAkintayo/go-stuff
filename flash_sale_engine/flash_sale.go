package main

import (
	"errors"
	"fmt"
	"time"
)

type Inventory struct {
	token int
}

func populateInventory(inventoryChan chan Inventory, totalUnits int) (bool, error) {
	unitsNeeded := 5
	if totalUnits != unitsNeeded || cap(inventoryChan) != unitsNeeded {
		errorMessage := fmt.Sprintf("Not exactly %d units", unitsNeeded)
		return false, errors.New(errorMessage)
	}
	for i := 0; i < cap(inventoryChan); i++ {
		newInventory := Inventory{token: i * 5}
		inventoryChan <- newInventory
	}
	return true, nil
}

// The monitor now listens directly to sales events instead of a blind clock ticker
func monitorInventory(salesChan <-chan int, totalUnits int, doneChan chan<- bool) {
	itemsSold := 0
	
	for range salesChan {
		itemsSold++
		remaining := totalUnits - itemsSold
		fmt.Printf("[DASHBOARD] Logged Sale! Phones Sold: %d | Remaining Stock: %d\n", itemsSold, remaining)
		
		if itemsSold == totalUnits {
			fmt.Println("[DASHBOARD] All stock cleared permanently! Signaling main thread to shut down.")
			doneChan <- true
			return
		}
	}
}

func RunFlashSaleEngine() {
	totalUnits := 5
	inventoryChan := make(chan Inventory, totalUnits)
	
	// Channels to coordinate the dashboard lifecycle
	salesChan := make(chan int, totalUnits)
	doneChan := make(chan bool)

	populateInventory(inventoryChan, totalUnits)

	// Start the event-driven monitor
	go monitorInventory(salesChan, totalUnits, doneChan)

	fmt.Println("Starting monitoring session. Flash sale is LIVE!")

	// Fixed the loop limit: exactly 50 customers (1 to 50)
	for i := 1; i <= 50; i++ {
		customerID := i
		go func() {
			unitPulled := <-inventoryChan
			
			customerTimeout := time.After(2 * time.Second)
			paymentDone := make(chan bool, 1) // Buffered to prevent memory leaks!

			go func() {
				time.Sleep(time.Duration(1000+((customerID*500)%2500)) * time.Millisecond)
				paymentDone <- true
			}()

			select {
			case <-paymentDone:
				fmt.Printf("Customer %d: Payment successful for token %d!\n", customerID, unitPulled.token)
				salesChan <- customerID // Notify monitor of a successful sale

			case <-customerTimeout:
				fmt.Printf("Customer %d: Payment TIMEOUT! Returning token %d to inventory.\n", customerID, unitPulled.token)
				inventoryChan <- unitPulled
			}
		}()
	}

	// The ultimate cleanup: Main thread blocks here until the monitor says stock is 0
	<-doneChan
	fmt.Println("Flash Sale Engine shut down cleanly. Everything closed!")
}