package main

import "fmt"	

func main() {
	 // IIFE: This function executes immediately
	func(){
		fmt.Println("Hello, World!")
	}() // <--- These parentheses trigger the execution

	// Passing Arguments Inline
	func(a, b int) {
		fmt.Println(a + b)
	}(3, 4) // a=3, b=4 


	// Assigning to a Variable
	result:=func(a,b int)int{
		return a*b
	}

	fmt.Println(result(3, 4))
	
}
