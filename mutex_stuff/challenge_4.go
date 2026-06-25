/*
Challenge 4: The Circular Waiting Trap (Very Hard / Deadlocks)
Now, let's step into the danger zone. In large-scale architecture, the biggest threat when dealing with multiple locks isn't data corruption—it's a Deadlock.

A deadlock occurs when two or more goroutines are frozen forever because they are stuck in a circular waiting chain.

Your Task
Write a function where you create two independent structs (e.g., WalletA and WalletB), each containing its own balance and its own sync.Mutex.

Simulate a transaction where money needs to be transferred between them concurrently (e.g., Worker 1 tries to transfer from A to B, while Worker 2 tries to transfer from B to A at the exact same millisecond).

Deliberately construct the key acquisition sequence so that your program freezes completely (deadlocks) when run.

Then, show me the fixed version and explain the design pattern you used to break the circle.
*/

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
	"unsafe"
)

type NewWallet struct {
	balance float64
	key     sync.Mutex
}

// Deposit adds the specified amount to the wallet balance.
func (nw *NewWallet) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}
	nw.key.Lock()
	defer nw.key.Unlock()
	nw.balance += amount
	return nil
}

// Withdraw subtracts the specified amount from the wallet balance after validation.
func (nw *NewWallet) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be greater than zero")
	}
	nw.key.Lock()
	defer nw.key.Unlock()
	if nw.balance < amount {
		return errors.New("insufficient balance")
	}
	nw.balance -= amount
	return nil
}

// GetBalance returns the current balance of the wallet safely.
func (nw *NewWallet) GetBalance() float64 {
	nw.key.Lock()
	defer nw.key.Unlock()
	return nw.balance
}

// transferDeadlock deliberately constructs a circular waiting chain causing a deadlock.
// Worker 1 locks fromWallet, Worker 2 locks toWallet, and then both block forever waiting for the other.
func transferDeadlock(fromWallet *NewWallet, toWallet *NewWallet, amount float64) {
	fromWallet.key.Lock()
	time.Sleep(10 * time.Millisecond) // Force a context switch to guarantee deadlock triggers
	toWallet.key.Lock()

	fromWallet.balance -= amount
	toWallet.balance += amount

	toWallet.key.Unlock()
	fromWallet.key.Unlock()
}

// Preventing deadlock using Lock Ordering (Pattern 1)
// We compare memory addresses using unsafe.Pointer to enforce a consistent locking order across all goroutines.
func transferFromA(fromWallet *NewWallet, toWallet *NewWallet, amount float64) error {
	if amount <= 0 {
		return errors.New("transfer amount must be greater than zero")
	}
	if fromWallet == toWallet {
		return errors.New("cannot transfer to the same wallet")
	}

	// Always acquire locks in order of increasing memory address to break circular dependency
	if uintptr(unsafe.Pointer(fromWallet)) < uintptr(unsafe.Pointer(toWallet)) {
		fromWallet.key.Lock()
		toWallet.key.Lock()
	} else {
		toWallet.key.Lock()
		fromWallet.key.Lock()
	}
	defer fromWallet.key.Unlock()
	defer toWallet.key.Unlock()

	// Verify balance inside the locked critical section
	if fromWallet.balance < amount {
		return errors.New("insufficient balance for transfer")
	}

	fromWallet.balance -= amount
	toWallet.balance += amount
	return nil
}

// Prevent deadlock using TryLock/backoff (Pattern 2)
// If we can't acquire the second lock immediately, we release the first lock and try again later.
func transferFromB(fromWallet *NewWallet, toWallet *NewWallet, amount float64) error {
	if amount <= 0 {
		return errors.New("transfer amount must be greater than zero")
	}
	if fromWallet == toWallet {
		return errors.New("cannot transfer to the same wallet")
	}

	for {
		fromWallet.key.Lock()

		// Optimize: Check balance before trying to lock the destination wallet
		if fromWallet.balance < amount {
			fromWallet.key.Unlock()
			return errors.New("insufficient balance for transfer")
		}

		// Try to lock the second wallet. If successful, perform transfer and release both locks.
		if toWallet.key.TryLock() {
			fromWallet.balance -= amount
			toWallet.balance += amount

			toWallet.key.Unlock()
			fromWallet.key.Unlock()
			return nil
		}

		// If we couldn't lock toWallet, unlock fromWallet to allow other goroutines to use it,
		// and sleep (backoff) to prevent CPU hogging.
		fromWallet.key.Unlock()

		backOff := time.Duration(rand.Intn(10)+1) * time.Millisecond
		time.Sleep(backOff)
	}
}

func Challenge4() {
	walletA := NewWallet{balance: 100_000.0}
	walletB := NewWallet{balance: 150_000.0}

	fmt.Println("--- Starting safe concurrent transfers (Lock Ordering) ---")
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := transferFromA(&walletA, &walletB, 20_000.0); err != nil {
			fmt.Printf("[Transfer A->B] Error: %v\n", err)
		} else {
			fmt.Println("[Transfer A->B] Successful.")
		}
	}()

	go func() {
		defer wg.Done()
		if err := transferFromA(&walletB, &walletA, 30_000.0); err != nil {
			fmt.Printf("[Transfer B->A] Error: %v\n", err)
		} else {
			fmt.Println("[Transfer B->A] Successful.")
		}
	}()

	wg.Wait()
	fmt.Printf("Final Balances after safe transfer:\n  Wallet A: %.2f\n  Wallet B: %.2f\n\n", walletA.GetBalance(), walletB.GetBalance())

	// Note: If you want to witness a deadlock freeze, uncomment the following code:
	
		fmt.Println("--- Starting deadlocking concurrent transfers. This is expected to freeze ---")
		go transferDeadlock(&walletA, &walletB, 10_000.0)
		go transferDeadlock(&walletB, &walletA, 10_000.0)
		time.Sleep(1 * time.Second) // wait to show it's frozen
	
}
