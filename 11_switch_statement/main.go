package main

import "fmt"

func main() {

	// Go automatically breaks out of switch cases, so you do not need to write break.
	day:=4
	switch day{
		case 1:
			fmt.Println("Monday")
		case 2:
			fmt.Println("Tuesday")
		case 3:
			fmt.Println("Wednesday")
		case 4:
			fmt.Println("Thursday")
		case 5:
			fmt.Println("Friday")
		case 6:
			fmt.Println("Saturday")
		case 7:
			fmt.Println("Sunday")
		default:
			fmt.Println("Invalid day")
	}

	marks:=250
	switch {
		case marks>=550:
		   fmt.Println("Excellent")
		case marks>=450:
			fmt.Println("Good")
		case marks>=350:
			fmt.Println("Average")
		case marks>=250:
			fmt.Println("Below Average")
		default:
			fmt.Println("Invalid marks")
	}
}
