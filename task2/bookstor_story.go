package task2

import "fmt"

type Car struct {
	Make string
	Model string
	RentPerDay float64
}
 func (c Car) CalculateRent(days int) float64 {
	return c.RentPerDay * float64(days)
 }

 func (c Car) Display(){
	fmt.Println("Make:", c.Make)
	fmt.Println("Model:", c.Model)
	fmt.Println("Rent Per Day:", c.RentPerDay)
 }