package main

import "fmt"

type InvalidAmountError struct {
	amount float64
}

type InsufficientFundsError struct {
	balance   float64
	requested float64
}

func (e InvalidAmountError) Error() string {
	return fmt.Sprintf(
		"cannot purchase airtime of N%.2f. Amount must be greater than zero",
		e.amount)
}

func (e InsufficientFundsError) Error() string {
	return fmt.Sprintf("insufficient funds: available balance is N%.2f, but you requested N%.2f", 
	e.balance, 
	e.requested)
}

func BuyAirtime(balance float64, amount float64) (float64, error) {
	if amount <= 0 {
		return balance, InvalidAmountError{amount}
	} 
	if amount > balance {
		return balance, InsufficientFundsError {balance, amount}
	}

	newBalance := balance - amount
	return newBalance, nil
}

func main(){
	balance := 2000.0

	newBalance, err := BuyAirtime(balance, 5000.0)

	if (err != nil){
		fmt.Println("Too huge amount: ", err)
	} else {
		balance = newBalance
		fmt.Printf("Airtime purchase successful New balance is %.2f", newBalance)
	}

	newBalance, err = BuyAirtime(balance, -50.0)
	if (err != nil){
		fmt.Println("Maths Error: ", err)
	} else {
		balance = newBalance
		fmt.Printf("Airtime purchase successful New balance is %.2f", newBalance)
	}

	newBalance, err = BuyAirtime(balance, 500.0)
	if (err != nil){
		fmt.Println(err)
	} else {
		balance = newBalance
		fmt.Printf("Airtime purchase successful. New balance is %.2f", newBalance)
	}
}
