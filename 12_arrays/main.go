package main

import "fmt"

func main(){
	// arrays in go are fixed and cannot grow
	var marks[3]int
	marks[0]=10
	marks[1]=11
	marks[2]=12

	fmt.Println(marks)

	// array literal
	res:=[5]int{2,3,4,5,6}
	fmt.Println(res)
	fmt.Println(len(res))
}