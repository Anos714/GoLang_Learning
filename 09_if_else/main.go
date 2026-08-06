package main

import (
	"fmt"
)

func main() {
	score:=72
	if score>=80{
		fmt.Println("you achieve a high score")
	}else{
		fmt.Println("you need to improve")
	}

	age:=12
	if age<18{
		fmt.Println("you are a minor")
	}else{
		fmt.Println("you are an adult")
	}

	op:="+"
	if(op=="+"){
		fmt.Println("addition")
	}else if(op=="-"){
		fmt.Println("subtraction")
	}else{
		fmt.Println("invalid operator")
	}

	// if with short statement
	items:=3;
	pricePerItem:=49

	if total:=items*pricePerItem;total>=100{
		fmt.Println("total is bohot jyada", total)
	}else{
		fmt.Println("total is thoda kam", total)
	}
}
