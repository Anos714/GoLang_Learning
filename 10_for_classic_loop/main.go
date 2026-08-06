package main

import (
	"fmt"
)

func main(){
	// classic for loop
	for i:=0; i<10;i++{
		fmt.Println("i: ",i)
	}

	// add a numnber n times
	n:=10
	sum:=0
	for i:=1;i<=n;i++{
		sum+=i
	}
	fmt.Println("sum: ",sum)
}