package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
)

type Wallet struct {
	balance float64
	lock    sync.RWMutex
}

func (wallet *Wallet) GetBalance() float64 {
	wallet.lock.RLock()
	defer wallet.lock.RUnlock()
	return wallet.balance
}

func (wallet *Wallet) Deposit(amount float64) bool {
	wallet.lock.Lock()
	defer wallet.lock.Unlock()
	wallet.balance += amount
	return true
}

func (wallet *Wallet) Withdraw(amount float64) (bool, error) {
	wallet.lock.Lock()
	defer wallet.lock.Unlock()
	if amount > wallet.balance {
		return false, errors.New("Insufficient balance")
	}
	wallet.balance -= amount
	return true, nil
}

func collector(collectorChan chan int, doneChan chan bool) {
	var balanceCalls, depositCalls, withdrawCalls int

	for count := range collectorChan {
		switch count {
		case 0:
			balanceCalls++
		case 1:
			depositCalls++
		case 2:
			withdrawCalls++
		}
	}

	fmt.Printf("\n\nBalance calls: %d\nDeposit calls: %d\nWithdraw Calls: %d", balanceCalls, depositCalls, withdrawCalls)
	doneChan <- true
}

func HighTraffic() {
	newWallet := Wallet{balance: 100_000.0}
	amountToDeposit := 20.0
	amountToWithdraw := 10.0
	collectorChan := make(chan int, 1000)
	doneChan := make(chan bool)

	// Spin up 1000 goroutines and in each loop they will randomly pick between any of the three functions;
	var routineTracker sync.WaitGroup

	go collector(collectorChan, doneChan)

	for i := 1; i <= 1000; i++ {
		routineTracker.Add(1)

		go func() {
			defer routineTracker.Done()
			switch rand.Intn(3) {
			case 0:
				balance := newWallet.GetBalance()
				collectorChan <- 0
				fmt.Printf("Balance is %f\n", balance)

			case 1:
				newWallet.Deposit(amountToDeposit)
				collectorChan <- 1
				fmt.Printf("User deposited %f\n", amountToDeposit)

			case 2:
				newWallet.Withdraw(amountToWithdraw)
				collectorChan <- 2
				fmt.Printf("User just withdrew %f\n", amountToWithdraw)
			}
		}()
	}

	routineTracker.Wait()
	close(collectorChan)
	<- doneChan
}
