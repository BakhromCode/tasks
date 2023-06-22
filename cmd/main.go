package main

import (
	"fmt"
	"github/tasks/task3"
)

func main() {

	bankAccount := task3.BankAccount2{
		AccountNumber: "0001",
		HolderName:    "Bek",
		Balance:       1000,
		Transaction:   []task3.Transaction{},
	}

	bankAccount.Deposit(200)
	bankAccount.Withdraw(300)
	bankAccount.Display()

	/////////////////////////////////////////
	// Lek := task2.Social{
	// 	Name: "Bek",
	// 	Email: "Bek@mail.com",
	// 	Friends: []task2.Social{},
	// }
	// Jek := task2.Social{
	// 	Name: "Bek",
	// 	Email: "Bek@mail.com",
	// 	Friends: []task2.Social{},
	// }
	// Modrin := task2.Social{
	// 	Name: "Bek",
	// 	Email: "Bek@mail.com",
	// 	Friends: []task2.Social{},
	// }

	// Lek.AddFriends(Modrin) {
	// 	Name: "Bek",
	// 	Email: "Bek@mail.com",
	// 	Friends: []task2.Social{},
	// }

	//////////////////////////////////////////
	// car := task2.Car {
	// 	Make: "Toyata",
	// 	Model: "Corolla",
	// 	RentPerDay: 40.0,
	// }

	// fmt.Println("10 kunga: ", car.CalculateRent(14), "$")
	// car.Display()

	// account := task2.BankAccount{
	// 	AccountNumber: "123456789",
	// 	HolderName:    "John Doe",
	// 	Balance:       1000.0,
	// }

	// account.Deposit(500.0)
	// account.Withdraw(200.0)

	// account.Display()

	// fmt.Println((task1.Nma(1)))
	// fmt.Println((task1.Nma(2)))
	// fmt.Println((task1.Nma(3)))
	// fmt.Println((task1.Nma(4)))
	// fmt.Println((task1.Nma(-1)))
	// fmt.Println((task1.Nma(-2)))
	// fmt.Println((task1.Nma(-3)))

	// print("Hello world")

	// var arr [] int = [] int{ 1, 2, 5, 8, 9, 6, 4 }

	// toq, juft := task1.OddEvenIndex(array)
	// fmt.Printf(toq,juft)

}
