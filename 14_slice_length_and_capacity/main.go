package main

import "fmt"

func main() {

/*
// 1. Literal declaration (Best when you know the initial values)
languages := []string{"Go", "Python", "SQL"}

// 2. Using make() (Best when you know the expected size upfront - highly optimized)
// make(Type, Length, Capacity)
users := make([]User, 0, 100) // Space allocated for 100 users, initial length 0

// 3. Slicing an existing array or slice
numbers := [5]int{10, 20, 30, 40, 50}
subset := numbers[1:4] // Contains: [20, 30, 40]
 */


	scores:=make([]int,0,5)
	fmt.Println(scores,len(scores),cap(scores))


	// if in case we r excedding the capacity, the slice will be automatically resized(or in other words go grows the backing array in double size)
	scores = append(scores, 10, 20, 30, 40, 50,60,70)
	fmt.Println(scores)

	// slicing the slice
	fmt.Println(scores[1:5])


	todos:=[]string{"learn Go","learn hono","build atleast one project in go and hono","try to secure internships","learn ai and devops"}
	fmt.Println(todos,len(todos),cap(todos))

	moreTodos:=[]string{"also do workout if possible","please be productive"}

	//  passing arguments to variadic parameters or variadic argument expansion or spread operator (...)
	todos=append(todos, moreTodos...)
	fmt.Println(todos,len(todos),cap(todos))
}



