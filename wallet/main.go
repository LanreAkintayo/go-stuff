package main

import "fmt"


type Wallet struct {
	balance float64
}

func (w *Wallet) Deposit(amount float64) {
	w.balance += amount
}

func (w Wallet) GetBalance() float64 {
	return w.balance
}

func main() {
	wallet := Wallet{
		balance: 0.0,
	}

	amount := 5000.0
	wallet.Deposit(amount)

	newBalance := wallet.GetBalance()

	fmt.Printf("The balance is %.2f", newBalance)
}