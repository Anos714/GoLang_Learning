package main

import (
"fmt"
"strings"
)



func main(){
	firstName:="Rahul"
	LastName:="Sain"

	fullName:=firstName+" "+LastName
	fmt.Println("Full Name: ",fullName)

	fmt.Println(strings.ToUpper(fullName))
	fmt.Println(strings.ToLower(fullName))
}