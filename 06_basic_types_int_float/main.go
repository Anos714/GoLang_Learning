package main

import (
	"fmt"
)

func main(){
	views1:=1000
	views2:=2000
	totalViews:=views1+views2
	fmt.Println("total views: ",totalViews)

	views1++
	fmt.Println("views1 increment: ",views1)
	views1--
	fmt.Println("views1 decrement: ",views1)

	avgViews:=totalViews/2
	fmt.Println("Average views: ",avgViews)

	weight1:=68.800
	weight2:=70.700

	combinedWeight:=weight1+weight2
	avgWeight:=combinedWeight/2

	fmt.Println("combined weight: ",combinedWeight)
	fmt.Println("avg weight: ",avgWeight)

}
