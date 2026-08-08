package main

import "fmt"


type Student struct{
	Roll_No uint
	Name string
	Age uint
	Class uint
}

func (student *Student) printStudentName() string{
	fmt.Println("student address: ",&student)
	fmt.Println("student: ",*student)
	*&student.Name="Bajaj singh"
	*&student.Age=16
	return fmt.Sprintf("Hi i am %s",student.Name)
}

func main(){
	student:=Student{
		Roll_No: 1,
		Name: "Bajaj",
		Age: 15,
		Class: 10,
	}

	res:=student.printStudentName()
	fmt.Println("Res: ",res)
	fmt.Println("After student struct: ",student)
}
