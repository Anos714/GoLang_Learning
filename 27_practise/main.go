package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)



func main(){
	// 1. variables
	// 1st way
	var age int
	age=21
	fmt.Println(age)

	// 2nd way
	// we dont give name a type because it inference it
	var name="Rahul Sain"
	fmt.Println(name)

	// 3rd way
	gender:="male"
	fmt.Println(gender)

	// 4th way
	likes,dislikes,views,comments,shares :=1000,2000,7000,500,100
	fmt.Printf("likes: %v\ndislikes: %v\nviews: %v\ncomments: %v\nshares: %v\n",likes,dislikes,views,comments,shares)
// -----------------------------------------------------------------------

	// package imports (math and strings pkg)
	fmt.Println(math.Pi)
	fmt.Println(math.Pow(20,2))
	fmt.Println(math.Pow10(7))
	fmt.Println(math.Floor(22.23))
	fmt.Println(math.Ceil(22.33))
	fmt.Println(math.Round(22.23))
	fmt.Println(math.IsNaN(2345))
	fmt.Println(rand.Intn(100)+1)


	fmt.Println(strings.ToLower("HELLEO"))
	fmt.Println(strings.ToTitle("hello"))
	fmt.Println(strings.ToUpper("hello"))
	fmt.Println(strings.Contains("hello","llo"))
	str:="rahul sain"
	str1:=strings.Clone(str)
	str="bhai"
	fmt.Println(str) //bhai
	fmt.Println(str1) // rahul sain


	// -----------------------------------------------------------------------

	/*
	Go Data Types Hierarchy
	├── 1. Basic (Primitive) Types
	│   ├── Booleans
	│   │   └── bool (true/false, no implicit conversion)
	│   │
	│   ├── Numeric Types
	│   │   ├── Signed Integers (Can be negative)
	│   │   │   ├── int8, int16, int32, int64
	│   │   │   └── int (Platform-dependent: 32 or 64-bit)
	│   │   │
	│   │   ├── Unsigned Integers (Positive & zero only)
	│   │   │   ├── uint8, uint16, uint32, uint64
	│   │   │   └── uint (Platform-dependent: 32 or 64-bit)
	│   │   │
	│   │   ├── Special Integer Aliases
	│   │   │   ├── byte (alias for uint8, raw data/ASCII)
	│   │   │   └── rune (alias for int32, Unicode characters)
	│   │   │
	│   │   ├── Floating-Point (Decimals)
	│   │   │   ├── float32
	│   │   │   └── float64 (Go's default decimal type)
	│   │   │
	│   │   └── Complex Numbers (Advanced math)
	│   │       ├── complex64
	│   │       └── complex128
	│   │
	│   └── Strings
	│       └── string (Immutable sequence of UTF-8 bytes)
	│
	├── 2. Aggregate (Composite) Types
	│   ├── Array  (Fixed-size sequence of same type, e.g., [5]int)
	│   └── Struct (Collection of named fields, e.g., type User struct)
	│
	├── 3. Reference Types (Pointers & Dynamic Collections)
	│   ├── Pointer   (Holds memory address of a value, e.g., *int)
	│   ├── Slice     (Dynamically-sized view into an array, e.g., []int)
	│   ├── Map       (Key-value pairs, unordered, e.g., map[string]int)
	│   ├── Channel   (Used to pass data between Go routines safely)
	│   └── Function  (Functions are first-class types in Go)
	│
	└── 4. Interface Type
    └── interface{} / any (Defines method sets; can hold any value)
	*/


	// to check type we use %T or reflect pkg
	fmt.Printf("%T\n",20) // int
	fmt.Printf("%T\n",20.22) // float64
	fmt.Printf("%T\n",20i-12) // complex128
	fmt.Printf("%T\n",true) // bool
	fmt.Printf("%T\n","rahul") // string
	 arr:=[3]int{1,2,3}
	fmt.Printf("%T\n",arr) // [3]int (array)
	slc:=[]int{1,2,3}
	fmt.Printf("%T\n",slc) // []int (slice)

	type User struct{
		Name string
	}
	var user User
	user.Name="Hello"
	fmt.Printf("%T\n",user) // main.User (struct)

	map1:=map[string]int{
		"rahul":21,
		"nitin":21,
	}
	fmt.Printf("%T\n",map1) // map[string]int

	var a=20
	var b *int=&a
	fmt.Printf("%T\n",b) // *int (pointer)


	// ---------------------------------------------

	// constants -> constants are values that are fixed at compile-time and cannot be changed during the execution of your program. They are incredibly efficient because the Go compiler replaces the constant name directly with its literal value wherever it is used, causing zero runtime memory overhead.
	const MAX_VALUE=2500
	const MAX_TO_UPLOAD=25


	// ---------------------------------------------

	// conditionals
	// 1. if-else
	isVoter:=true
	if !isVoter{
		fmt.Println("You cannot vote, you are underage")
	}else{
		fmt.Println("You can vote")
	}

	// 2. for loops
	// a) classic for loop
	for i:=1;i<=5;i++{
		fmt.Println(i)
	}

	// b) for range loop
	fruits:=[]string{"apple","mango","banana","guava"}
	for index,val:=range fruits{
		fmt.Printf("at index: %v -> %v\n",index,val)
	}

	// 3. switch statement
	score:=72
	switch {
		case score>=75:
		fmt.Println("Excellent")
		case score>=55:
		fmt.Println("Very good")
		case score>=35:
		fmt.Println("Good")
		case score<=34:
		fmt.Println("you can do better")
		default:
		fmt.Println("Enter the valid score")
	}

	// ---------------------------------------------

	// array(it is of fixed size) and slice(it is of dynamic size)

	// 1. array -> [length]type -> [3]int
	// 1st way
	var arr1[3]int
	// both ways you can put values in array or slices
	arr[0]=12
	arr[1]=23
	arr[2]=24

	arr1=[3]int{12,22,25}
	fmt.Println(arr1)

	// 2nd way
	arr2:=[3]int{2,3}
	fmt.Println(arr2)

	// length of array and capacity of array
	fmt.Println(len(arr2))
	fmt.Println(cap(arr2))

	// 2. slice -> []type -> []string
	// 1st method
	var slc1[]string
	slc1=[]string{"hi","hello","bye"}
	fmt.Println(slc1)

	// 2nd method
	slc2:=[]int{1,2,3,4,5,6,7}
	fmt.Println(slc2)

	// using make() to make slice in this which can define how much items present and the capacity also
	slc3:=make([]int,0,5) // items present -> 0, and capacity -> 5
	slc3 = append(slc3, 10)
	slc3=append(slc3, 20)
	slc3=append(slc3, 20)

	// length and capacity of slice
	fmt.Println(len(slc3))
	fmt.Println(cap(slc3))


		// ---------------------------------------------

		// map
		userMap:=map[string]any{
			"name":"Rahul sain",
			"age":21,
			"gender":"male",
			"canVote":true,
		}

		fmt.Println(userMap)

		// for range loop over map
		for key,val:=range userMap{
			fmt.Printf("at key: %v -> value: %v\n",key,val)
		}

		// map delete() property
		delete(userMap,"age")

		fmt.Println(userMap)
		fmt.Println(userMap["name"])


		// comma-ok idiom -> it checks that a key is present in the map or not
		if item,ok:=userMap["age"];ok{
			fmt.Println(item,ok)
		}else{
			fmt.Println(ok)
		}


// ----------------------------------------------------

	// pointers
	// Do use pointers for: Structs, Arrays, and Primitives (if mutating).
	// Don't use pointers for: Slices, Maps, Channels, and 'any' (Go handles references natively)
	fruit:="mango"
	fmt.Printf("before address: %p\n",&fruit)
	fmt.Println("Before Value: ",fruit)
	getData(&fruit)
	fmt.Println("after function called: ",fruit)
}


func getData(fruit *string){
	fmt.Printf("after address: %p\n",fruit)
	*fruit="apple"
	fmt.Println("after Value: ",*fruit)
}
