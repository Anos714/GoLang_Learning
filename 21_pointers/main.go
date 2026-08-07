package main

import "fmt"

func addScore(score *int){
	fmt.Println("Score address: ",score)
	fmt.Println("Score value: ",*score)
	*score = *score + 10
	fmt.Println("Updated score value: ",*score)
}

func salaryHikeByYear(salary *int,hikePerecent float64){
	fmt.Println("Salary Address:",salary)
	fmt.Println("Salary Value: ",*salary)

	hikeValue:=int(float64(*salary)*hikePerecent/100)
	*salary=*salary+hikeValue
	fmt.Println("Updated Salary Value: ",*salary)
}

func main(){
	// pointer -> store the memory address of any value

	// GoLang Address (&)
	// &x -> address of x (makes a pointer)

	// Dereferencing (*)
	// *p -> value at the memory address p (we can read write to it)

	score:=10
	fmt.Println("Before score: ",score)
	fmt.Println("Score address: ",&score)

	addScore(&score)
	fmt.Println("After Score: ",score)


	salary:=35000
	hikePercent:=13.50
	fmt.Println("Salary Before Hike: ",salary)
	fmt.Println("Salary Address: ",&salary)

	salaryHikeByYear(&salary,hikePercent)
	fmt.Println("Salary After Hike: ",salary)
}






/*
=============================================================================
GOLANG POINTERS DOCUMENTATION & CHEAT SHEET
=============================================================================

1. VARIABLES AND MEMORY ADDRESSES
-----------------------------------------------------------------------------
Every variable created is stored in a specific location in the computer's
memory. This location has a unique hexadecimal identifier called a "memory address".

The '&' Operator (Address-of):
- Used to retrieve the exact memory address of an existing variable.
- Example: &age will return something like 0xc0000120b8.
*/

/*
2. WHAT IS A POINTER?
-----------------------------------------------------------------------------
A pointer is a specialized variable that stores the memory address of another
variable, rather than storing a direct data value.

The Type Syntax (*T):
- The type of a pointer is written as '*T', where 'T' represents the data type
  of the value it points to (e.g., *int, *string, *float64).
- A pointer of type '*int' can ONLY hold the memory address of an integer variable.
*/

/*
3. DEREFERENCING (ACCESSING THE UNDERLYING VALUE)
-----------------------------------------------------------------------------
Once you have stored a memory address inside a pointer, you need a way to
read or modify the actual data residing at that address. This is called "dereferencing".

The '*' Operator (Dereference):
- Placing an asterisk ('*') before a pointer variable acts as a key to enter
  the underlying memory space.
- Read operation:  val := *p   (Reads the value the pointer is looking at)
- Write operation: *p = 30     (Overwrites the value directly in memory)
*/

/*
4. QUICK SYNTAX REFERENCE SUMMARY
-----------------------------------------------------------------------------
Syntax       | What it does                   | Real-World Analogy
-------------|--------------------------------|-------------------------------
age := 25    | Creates a data value           | Building a physical house.
&age         | Fetches the memory address     | Looking up the GPS coordinates.
p := &age    | Stores address in a pointer    | Writing GPS coordinates on a paper.
*p           | Dereferences the pointer       | Traveling to coordinates to edit inside.
*/

/*
5. CRUCIAL RULES & BEHAVIORS TO REMEMBER
-----------------------------------------------------------------------------
A. Zero Value is 'nil':
   If a pointer is declared without initialization, its default value is 'nil'
   (it points to absolutely nothing). Attempting to dereference a 'nil' pointer
   (*nilPointer) will instantly cause a runtime panic and crash the program.

B. No Pointer Arithmetic:
   Unlike languages like C or C++, Go completely forbids pointer arithmetic
   by default (e.g., you cannot do 'p++' or 'p + 2'). This structural restriction
   ensures Go's memory-safety model.

C. The new() Built-in Function:
   You can allocate memory for an anonymous, zero-valued variable using 'new(T)'.
   It allocates the memory structure, sets it to its default zero value, and
   returns a pointer (*T) pointing directly to it.
*/

/*
6. WHY USE POINTERS IN GO?
-----------------------------------------------------------------------------
Reason A: Mutating Arguments Across Function Boundaries
- Go is strictly a "pass-by-value" language. When passing a standard variable
  to a function, Go duplicates the data. Modifying it inside the function
  does not affect the original caller variable.
- Passing a pointer passes a copy of the *address*. This allows the function
  to mutate the original external data directly.

Reason B: Performance and Memory Efficiency
- If a struct contains massive amounts of data (e.g., large arrays or text fields),
  passing it directly forces Go to duplicate the entire byte payload.
- Passing a pointer copies a lightweight 8-byte address regardless of how massive
  the underlying struct actually is.
*/

/*
import "fmt"

func main() {
	// --- Practical Example 1: Basic Address and Dereferencing ---
	age := 25
	var p *int = &age // 'p' stores the memory address of 'age'

	fmt.Println("Value of age:", age)   // Outputs: 25
	fmt.Println("Address of age:", &age) // Outputs: 0xc000...
	fmt.Println("Pointer value:", p)     // Outputs: 0xc000... (same address)
	fmt.Println("Dereferenced:", *p)    // Outputs: 25

	*p = 30 // Mutating the value via dereference
	fmt.Println("New value of age:", age) // Outputs: 30

	// --- Practical Example 2: Passing to Functions ---
	score := 10

	modifyWithoutPointer(score)
	fmt.Println("Score after pass-by-value:", score) // Outputs: 10 (unaffected)

	modifyWithPointer(&score)
	fmt.Println("Score after pointer pass:", score)   // Outputs: 99 (mutated)
}

// Function receives a direct copy of the integer data
func modifyWithoutPointer(val int) {
	val = 99
}

// Function receives the memory address copy, enabling direct access
func modifyWithPointer(val *int) {
	*val = 99
}

 */
