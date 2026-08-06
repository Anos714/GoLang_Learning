package main

import "fmt"

func main(){
	// most common collection type - it is dynamic and it can grow

	// syntax -> []type{...values} or var name[]type
	// ex:- var s []int or s := []string{"Go", "Rust"}

	names:=[]string{"Rahul","Rajesh","Nitin","Hemant","Krishh","Manish","Mohit"}

	fmt.Println(names)
	fmt.Println(names[0],names[len(names)-1])

	names[1]="Raja"

	for i,val:=range names{
		fmt.Println("Index: ",i,"\n","Value: ",val)
	}

	// this is also a slice
	var nums[]int
	nums=append(nums,10)
	nums=append(nums,20)
	nums=append(nums,30,50,70)
	nums=append(nums,40)

	fmt.Println(nums)

	// appending the names slice into namesArr Array 
	var namesArr[]string
	namesArr=append(names)
	fmt.Println(namesArr)




}
