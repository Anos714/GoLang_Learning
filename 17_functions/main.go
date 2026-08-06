package main

import "fmt"



func add(num1,num2 int)int{
	return num1+num2
}

func showDetails(fullName string, age int, weight float64, adult bool,){
	fmt.Printf("FullName: %s\nAge: %d\nWeight: %.2f\nAdult: %t\n",fullName,age,weight,adult)
}

func binarySearch(arr []int, target int) int {
	start:=0
	end:=len(arr)-1

	for start<=end{
		mid:=start+(end-start)/2
		if arr[mid]==target{
			return mid
		}else if arr[mid]<target{
			start=mid+1
		}else{
			end=mid-1
		}
	}
	return -1
}

// multiple return type function
func sumAndProduct(num1, num2 int) (int, int) {
	sum:=num1+num2
	product:=num1*num2
	return sum,product
}

func areYouEligibleForVote(age int)(bool,string){
	if age>=18 {
		return true, "You are eligible for vote"
	}
	return false, "You are not eligible for vote"
}

// named return or naked return
func passOrFail(score int)(isPass bool, message string){
	if score>=50{
		isPass=true
		message="You passed"
	}else{
		isPass=false
		message="You failed"

	}
	return
}

func divide(a,b int) (result float64) {
	result=float64(a)/float64(b)
	return
}

// variadic function -> a function that accepts a variable number of arguments of the same type
func sumAll(numbers ...int)int{
	total:=0
	for _,val:=range numbers{
		total+=val
	}
	return total
}


func main() {
	result:=add(12,22)
	fmt.Println(result)
	showDetails("Jhon Doe",30,75.5,true)
	arr:=[]int{1,2,3,4,5,6,7,8,9,10}
	target:=9
	res:=binarySearch(arr,target)
	fmt.Printf("The target %d was found at index %d\n",target,res)
	sum,product:=sumAndProduct(5,3)
	fmt.Printf("Sum: %d, Product: %d\n",sum,product)

	isEligible, message:=areYouEligibleForVote(17)
	fmt.Printf("Eligible: %t, Message: %s\n",isEligible,message)

	isPass,msg:=passOrFail(25)
	fmt.Printf("Pass: %t, Message: %s\n",isPass,msg)

	res1:=divide(12,5)
	fmt.Println(res1)

	fmt.Println(sumAll(12,3,4,5,6,7,8,9))

	values:=[]int{1,2,3,4,5,6,7,8,9}
	fmt.Println(sumAll(values...))

}
