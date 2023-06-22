package task2

import "fmt"

type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

func (ba *BankAccount) Deposit(amout float64) {
	if ba.Balance <= ba.Balance {
		ba.Balance -= amout
	}

}

func (ba *BankAccount) Withdraw(amount float64) {

}

func (ba BankAccount) Display() {
	fmt.Println("Account Number:", ba.AccountNumber)
	fmt.Println("Holder Name:", ba.HolderName)
	fmt.Println("Balance:", ba.Balance)
}
