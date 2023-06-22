package task2

import "fmt"


type Item struct {
	Name string
	Price float64
}



type shopping struct {
	Name string
	Email string
	Hamyon []Item
}


 func (c *shopping)AddToCard(mahsilot Item) {
	c.Hamyon=append(c.Hamyon, mahsilot)
 }
 func(c*shopping)RemoveFromCart(ayirish Item){
	yangiSavatcha := []Item{}


	for _, v := range c.Hamyon {
		if v == mahsilot {
			continue
		}
		yangiSavatcha = append(yangiSavatcha, v)
	}
	c.Hamyon = yangiSavatcha
 }