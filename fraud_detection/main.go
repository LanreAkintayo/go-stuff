package main

import "fmt"

type Status int

const (
	Pending  Status = 0
	Approved Status = 1
	Blocked  Status = 2
)

type Transaction struct {
	id     string
	amount float64
	status Status
}

func checkFraudulent(currentTransaction Transaction) bool {
	if currentTransaction.amount > 500_000 {
		return true
	}
	if currentTransaction.status == Blocked {
		return true
	}
	return false
}

func ScanTransactions(transactions []Transaction) []Transaction {
	fraudTransactions := []Transaction{}
	for i := 0; i < len(transactions); i++ {
		currentTransaction := transactions[i]
		isFraudulent := checkFraudulent(currentTransaction)
		if isFraudulent {
			fraudTransactions = append(fraudTransactions, currentTransaction)
		}
	}

	return fraudTransactions
}


func GetTransactionIds(transactions []Transaction) []string {
	transactionIds := []string {}
	for i := 0; i < len(transactions); i++ {
		transactionId := transactions[i].id
		transactionIds = append(transactionIds, transactionId)	
	}

	return transactionIds
}

func main() {
	transactions := []Transaction{
		{id: "transaction1", amount: 100_000.0, status: Approved},
		{id: "transaction2", amount: 200_000.0, status: Blocked},
		{id: "transaction3", amount: 500_001.0, status: Approved},
		{id: "transaction4", amount: 550_000.0, status: Approved},
	}

	fraudTransactions := ScanTransactions(transactions)

	transactionIds := GetTransactionIds(fraudTransactions)

	fmt.Println("Fraud Transactions: ", transactionIds)
}
