package main

import (
	"fmt"
)

type User struct{
	Name string
	Age int
}

// value receiver means this method receives a copy of the user
func (user User) userPrint(){
	fmt.Println(user)
}

func main(){
	user:=User{
		Name:"Rahul Sain",
		Age:21,
	}

user.userPrint()
}
