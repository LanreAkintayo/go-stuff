package main

import "fmt"

type Transaction struct {
	id     string
	amount float64
	status string
}


func EvictBadTransaction(transaction *Transaction) {
	if transaction.status == "canceled" || transaction.status == "failed" {
		*transaction = Transaction{id: "", amount: 0.0, status: "evicted"}
	}
}

func main() {

	t1 := Transaction{id: "tx_01", amount: 2_500.0, status: "completed"}
	t2 := Transaction{id: "tx_02", amount: 15_000.0, status: "failed"}
	t3 := Transaction{id: "tx_03", amount: 450.0, status: "pending"}
	t4 := Transaction{id: "tx_04", amount: 9000.0, status: "canceled"}

	transactionPointers := []*Transaction{&t1, &t2, &t3, &t4}

	for _, transactionPointer := range transactionPointers {
		EvictBadTransaction(transactionPointer)
	}

	fmt.Println("t1: ", t1)
	fmt.Println("t2: ", t2)
	fmt.Println("t3: ", t3)
	fmt.Println("t4: ", t4)
}
