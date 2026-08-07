package main

import (
	"fmt"
	"log"
	"strconv"
)


func divide(a,b float64)(float64,error){
	if b==0{
		return 0, fmt.Errorf("division by zero")
	}
	return a/b,nil
}

func main() {
	// go don't use exceptions for normal failures
	// functions -> return errors as normal return values


	// syntax
	// val,err := something()
	// if err != nil {
	// 	handle error
	// 	return
	// }

if err:=run();err!=nil{
	log.Fatal(err)
}

	if val, err := checkAge(18); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Age: ", val)
	}


	if val,err:=divide(10,0);err!=nil{
		log.Fatal(err)
	} else {
		fmt.Println("Result: ", val)
	}
}




func run()error{
	input:="2"
	level,err:=parseLevel(input)

if err!=nil{
	return err
}
fmt.Println("Selected level: ",level)
return nil
}

func parseLevel(str string)(int,error){
	if str == "" {
		return 0, fmt.Errorf("empty string")
	}
	val,err:=strconv.Atoi(str)
	if err!=nil{
		return 0,fmt.Errorf("Level must be a number")
	}

	if val<1||val>5{
		return 0, fmt.Errorf("Level must be between 1 and 5")
	}

	return val,nil
}

func checkAge(age int)(int,error){
	if age<18{
		return 0,fmt.Errorf("Age must be at least 18")
	}
	return age,nil
}
