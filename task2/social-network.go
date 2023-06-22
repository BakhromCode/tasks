package task2

import "fmt"

type Social struct {  
Name string
Email string
Friends []Social
}

func (u *Social) AddFriends (friend Social){
	u.Friends = append(u.Friends, friend)
}

func (u *)