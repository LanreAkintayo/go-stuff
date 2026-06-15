package main

import "fmt"

type Transaction struct {
	username string
	amount   float64
	location string
}

type FraudGuard func(transaction Transaction) bool


func highAmountGuard(transaction Transaction) bool {
		if transaction.amount > 5_000_000.0 {
			return true
		}
		return false
}

func foreignLocationGuard(transaction Transaction) bool {
		if transaction.location == "Lagos" || transaction.location == "Abuja" {
			return false
		} 
		return true
	}

func FilterTransactions(transactions []Transaction, fraudGuard FraudGuard) []Transaction {
	filteredTransactions := make([]Transaction, 0, len(transactions))
	for _, transaction := range transactions {
		isGuarded := fraudGuard(transaction)
		if isGuarded {
			filteredTransactions = append(filteredTransactions, transaction)
		}
	}
	return filteredTransactions
}

func main() {
	transactions := []Transaction{
		{username: "Tunde", amount: 5_000.0, location: "Lagos"},
		{username: "Tunde", amount: 5_000.0, location: "Lagos"},
		{username: "Chidi", amount: 6000000.0, location: "Abuja"},
		{username: "Funmi", amount: 12000.0, location: "London"},
		{username: "Tunde", amount: 4500.0, location: "Lagos"},
		{username: "Osas", amount: 8000000.0, location: "Lagos"},
	}

	highAmountTransactions := FilterTransactions(transactions, highAmountGuard)
	foreignTransactions := FilterTransactions(transactions, foreignLocationGuard)

	fmt.Println("High Amount Transactions: ", highAmountTransactions)
	fmt.Println("Foreign Transactions: ", foreignTransactions)

}	
