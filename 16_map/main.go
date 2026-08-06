package main

import "fmt"

func main() {
	// map[keyType]valueType
	votersAge:=map[string]int{
		"rahul":21,
		"hemant":24,
		"ballu":52,
		"kallu":18,
	}
	fmt.Println(votersAge)
	fmt.Println(votersAge["rahul"])
	fmt.Println(votersAge["newton"]) //0
	fmt.Println(votersAge[""])  //0
	fmt.Println(len(votersAge))

	// make() property
	studentGrades:=make(map[string]int)
	fmt.Println(studentGrades)
	studentGrades["Rahul"]=20
	studentGrades["Hemant"]=22
	fmt.Println(studentGrades)

	// delete() property
	delete(studentGrades, "Rahul")
	fmt.Println(studentGrades)

	fmt.Println(len(studentGrades))


	// value ok pattern or comma-ok idiom for map in go

	grade,ok:=studentGrades["Rahul"]
	fmt.Println(grade,ok)

	// it checks the key is present or not in the map and returns the value and ok status else fallback to else part
	if grade,ok:=studentGrades["Hemant"];ok{
		fmt.Println(grade,ok)
	}else{
		fmt.Println("Rahul not found",grade,ok)
	}


	// for range loop on map
	total:=0
	for item,grade:=range studentGrades{
		fmt.Println("\nItem: ",item,"\nGrade: ",grade)
		total+=grade
	}
	fmt.Println("total: ",total)
}
