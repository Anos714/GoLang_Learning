package main

import "fmt"

func main(){
	isLogged:=true
	isAdmin:=false
	hasSubscription:=true

	fmt.Println(isLogged,isAdmin,hasSubscription)

	fmt.Println(isLogged && isAdmin)
	fmt.Println(isLogged || hasSubscription)
}