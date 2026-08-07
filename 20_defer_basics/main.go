package main

import (
	"errors"
	"fmt"
)

func main() {
	// the defer statement is used to schedule a function call to be executed after the surrounding function returns

		// here first prints "main function executed" and then "defer statement executed"
	// defer fmt.Println("defer statement executed")
	// fmt.Println("main function executed")




	fmt.Println("Case1: success")
	if err := doWork(true); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Case2: fail early")
	if err := doWork(false); err != nil {
		fmt.Println(err)
	}
}

/*
 * Output:
Case1: success
start: resource acquired
work: doing something important
work: done
cleanup: resource released
Case2: fail early
start: resource acquired
cleanup: resource released
something went wrong. i am returning early
 */

func doWork(success bool)error{
	fmt.Println("start: resource acquired")

	// this defer will guarantee that the cleanup message is printed even if the function returns early
	// it will run on every return path
	// success return path
	// failure return path - early return
	// cleanup return path
	defer fmt.Println("cleanup: resource released")
	if(!success) {
		return errors.New("something went wrong. i am returning early")
	}
	fmt.Println("work: doing something important")
	fmt.Println("work: done")
	return nil
}
