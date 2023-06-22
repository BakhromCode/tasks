package task3

import (
	//"flag"
	"fmt"
	"time"
)

type Transaction struct {
	Amount      float64
	Transaction time.Time
}

type BankAccount2 struct {
	AccountNumber string
	HolderName    string
	Balance       float64
	Transaction   []Transaction
}

func (b *BankAccount2) Deposit(income float64) {
	b.Balance = b.Balance + income
	newTransaction := Transaction{
		Amount:      income,
		Transaction: time.Now(),
	}
	fmt.Println("Muvaffaqiyatli tranzaksiya")

	b.Transaction = append(b.Transaction, newTransaction)
}

func (b *BankAccount2) Withdraw(outcome float64) {
	if outcome < b.Balance {
		b.Balance = b.Balance - outcome
		fmt.Println("Muvaffaqiyatli tranzaksiya!")

		newTransaction := Transaction{
			Amount:      -outcome,
			Transaction: time.Now(),
		}
		b.Transaction = append(b.Transaction, newTransaction)
	} else {
		fmt.Println("Insuficient money!")
	}
}
func (b BankAccount2) Display() {
	var positive, negative []Transaction
	for _, v := range b.Transaction {
		if v.Amount < 0 {
			negative = append(negative, v)
		} else {
			positive = append(positive, v)
		}
	}
	fmt.Println("Hisob raqami: ", b.AccountNumber)
	fmt.Println("Egasi: ", b.HolderName)
	fmt.Println("Miqdori; ", b.Balance)
	fmt.Println("O'tkazmalar: ")
	fmt.Println(positive, negative)
}
