package main

import "fmt"

type InvalidAccountId struct {
	accountId string
}

type InsufficientBalance struct {
	accountId string
	amount    float64
	balance   float64
}

func (e InvalidAccountId) Error() string {
	return fmt.Sprintf("Account with ID %s is invalid", e.accountId)
}
func (e InsufficientBalance) Error() string {
	return fmt.Sprintf("Account with ID %s has balance %.2f but purchasing electricity worth %.2f", e.accountId, e.balance, e.amount)
}

func debitBankCard(accountId string, balance float64, amount float64) (float64, error) {
	if accountId == "" {
		return 0.0, InvalidAccountId{accountId}
	}
	if balance < amount {
		return 0.0, InsufficientBalance{accountId, amount, balance}
	}

	newBalance := balance - amount
	return newBalance, nil
}

func requestTokenFromDisco(accountId string) (string, error) {
	return "123456", nil
}

func BuyElectrcityToken(accountId string, balance float64, amount float64) (string, float64, error) {
	updatedBalance, err := debitBankCard(accountId, balance, amount)

	if err != nil {
		return "", balance, err
	}

	newToken, err := requestTokenFromDisco(accountId)

	if err != nil {
		return "", balance, err
	}

	return newToken, updatedBalance, nil
}

func main() {
	// User Balance: #10,000
	userBalance := 10_000.0
	userAccountId := "userA"
	amountToPurchase := 500.0

	generatedToken, updatedBalance, err := BuyElectrcityToken(userAccountId, userBalance, amountToPurchase)

	if err != nil {
		switch e := err.(type) {
		case InsufficientBalance:
			fmt.Printf("Top up needed for account %s\n", e.accountId)

		case InvalidAccountId:
			fmt.Printf("Wrong account id: %s", e.accountId)

		default:
			fmt.Println("Error occured: ", err)
		}
		return
	}

	userBalance = updatedBalance
	fmt.Printf("%s has successfully purchase electricity. \nToken is %s. \nRemaining Balance: %.2f\n", userAccountId, generatedToken, userBalance)

}
