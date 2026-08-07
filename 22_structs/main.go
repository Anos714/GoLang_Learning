package main

import "fmt"

// Defining a custom Struct blueprint
type Employee struct {
    Name   string
    Age    int
    Salary float64
    Active bool
}

type User struct{
	Id uint
	Username string
	Age uint
	Gender string
	Email string
	Password string
	Role string
}


func main() {
	// In Go (Golang), a struct (short for structure) is a user-defined type that allows you to group and bundle related fields of different data types into a single, cohesive unit


	// Approach A: Field Names initialization (Recommended)
    emp1 := Employee{
        Name:   "Rahul Sain",
        Age:    25,
        Salary: 35000.0,
        Active: true,
    }

    // Approach B: Positional initialization (Order must match exactly)
    emp2 := Employee{"Amit", 28, 45000.0, false}

    // Accessing individual fields using dot (.) notation
    fmt.Println(emp1.Name) // Outputs: Rahul Sain

    // struct field are mutable by default
    emp2.Salary = 50000.0   // Overwriting values

    fmt.Println("emp1: ",emp1)
    fmt.Println("emp2: ",emp2)

    // User struct
    user1:=User{
    Id: 1,
    Username: "Anos",
    Age: 21,
    Gender: "Male",
    Email: "anos@gmail.com",
    Password: "anos@gg123",
    Role: "user",
    }

    fmt.Println("user1: ",user1)
    fmt.Println("user1 username: ",user1.Username)

    user2:=User{2, "Amit", 28, "Male", "amit@gmail.com", "amit@gg123", "user"}
    fmt.Println("user2: ",user2)
    fmt.Println("user2 email: ",user2.Email)

    // partial user
    user3:=User{Username:"kalia"}
    fmt.Println("user3: ",user3) //user3:  {0 kalia 0    }
    fmt.Println("user3 username: ",user3.Username) //user3 username:  kalia
    fmt.Println("user3 email: ",user3.Email) //user3 email: ""

}




/*
=============================================================================
GOLANG STRUCTS DOCUMENTATION & CHEAT SHEET
=============================================================================

1. DECLARING AND INITIALIZING STRUCTS
-----------------------------------------------------------------------------
A struct (structure) is a user-defined type that lets you group and bundle
related fields of different data types into a single, cohesive unit. Go does
not have traditional object-oriented classes; structs are used instead.

Key Initialization Approaches:
- Named Fields: Explicitly mapping values to field names. (Highly Recommended)
- Positional Values: Passing values without field names. Order must match the
  struct definition exactly.
- Dot (.) Notation: Used to access or overwrite individual properties.
*/

/*
2. ZERO VALUES OF A STRUCT
-----------------------------------------------------------------------------
If you declare a struct variable without assigning explicit values, Go
automatically sets all internal fields to their respective default zero values:
- string  -> ""
- numbers -> 0 / 0.0
- boolean -> false
- pointer -> nil
*/

/*
3. STRUCT POINTERS AND AUTOMATIC DEREFERENCING
-----------------------------------------------------------------------------
When passing a large struct to a function, copying the entire data chunk wastes
memory. Passing a pointer (*T) copies a lightweight 8-byte memory address.

Automatic Dereferencing:
- In languages like C, you must write (*ptr).Field to access data via a pointer.
- Go handles this automatically. You can write ptr.Field directly, and Go
  dereferences it under the hood.
*/

/*
4. STRUCT EMBEDDING (COMPOSITION OVER INHERITANCE)
-----------------------------------------------------------------------------
Go completely rejects hierarchical class inheritance ("extends"). Instead, you
build complex data structures by nesting or embedding smaller structs inside
larger ones.

Field Promotion:
- When an inner struct is embedded anonymously (without an explicit field name),
  its fields are "promoted" directly to the outer struct for easy access.
*/

/*
5. STRUCT TAGS (CRUCIAL FOR REST & JSON APIs)
-----------------------------------------------------------------------------
When building backend APIs (e.g., using Go Fiber), you need to translate
PascalCase Go struct fields into standard camelCase JSON keys used by frontends.

Metadata Annotations:
- Written as backticks `json:"key_name"` next to fields.
- 'omitempty' can be added to omit/hide a field from the JSON output if it is
  left blank or holds its default zero value.
*/

/*
6. STRUCT METHOD RECEIVERS
-----------------------------------------------------------------------------
You can bind custom functions and operations directly to structs using
Method Receivers. This mimics how methods operate on class objects.

- Value Receiver (e T): Operates on an isolated copy of the data. Modifying
  data inside the method does not affect the original variable.
- Pointer Receiver (e *T): Operates on the original memory space. Allows the
  method to mutate fields permanently.
*/

/*
7. ANONYMOUS STRUCTS
-----------------------------------------------------------------------------
If you need an ad-hoc, throwaway data structure for a single localized task
(like formatting a single API response or writing a quick test payload),
you can declare a struct inline without creating an explicit 'type' definition.
*/


/*
import "fmt"

// --- 1 & 2 Blueprint: Basic Struct ---
type Employee struct {
	Name   string
	Age    int
	Salary float64
	Active bool
}

// --- 4 Blueprint: Struct Embedding ---
type ContactInfo struct {
	Email string
	Phone string
}

type Manager struct {
	Name        string
	Role        string
	ContactInfo // Anonymous Field Embedding (Enables field promotion)
}

// --- 5 Blueprint: Struct Tags for JSON ---
type User struct {
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"` // Hides field if empty
}

func main() {
	// --- Practical Example 1: Basic Initialization & Zero Values ---
	emp1 := Employee{
		Name:   "Rahul Sain",
		Age:    25,
		Salary: 35000.0,
		Active: true,
	}

	var blank Employee // Demonstrating zero values

	fmt.Println("--- Basic Struct ---")
	fmt.Printf("Employee 1: %+v\n", emp1)
	fmt.Printf("Blank Employee (Zero Values): %+v\n\n", blank)

	// --- Practical Example 2: Struct Pointers ---
	empPtr := &Employee{Name: "Priya", Age: 24}
	empPtr.Age = 25 // Automatic dereferencing under the hood

	fmt.Println("--- Struct Pointers ---")
	fmt.Printf("Pointer access name: %s, age: %d\n\n", empPtr.Name, empPtr.Age)

	// --- Practical Example 3: Composition & Embedding ---
	mgr := Manager{Name: "Vikram", Role: "HR"}
	mgr.Email = "hr@company.com" // Field promotion in action

	fmt.Println("--- Embedded Struct ---")
	fmt.Printf("Manager Email (Promoted): %s\n", mgr.Email)
	fmt.Printf("Manager Email (Explicit): %s\n\n", mgr.ContactInfo.Email)

	// --- Practical Example 4: Method Execution ---
	fmt.Println("--- Method Receivers ---")
	emp1.DisplayProfile() // Calls Value Receiver

	emp1.GiveHike(13.50)  // Calls Pointer Receiver (Changes original value)
	fmt.Printf("Salary after 13.5%% hike: %.2f\n\n", emp1.Salary)

	// --- Practical Example 5: Anonymous Structs ---
	config := struct {
		Port string
		Env  string
	}{
		Port: ":8080",
		Env:  "production",
	}

	fmt.Println("--- Anonymous Struct ---")
	fmt.Printf("Running on port %s in %s environment\n", config.Port, config.Env)
}

// --- 6 Implementation: Value Receiver (Reads data copy) ---
func (e Employee) DisplayProfile() {
	fmt.Printf("Profile: %s holds a salary baseline.\n", e.Name)
}

// --- 6 Implementation: Pointer Receiver (Mutates original memory) ---
func (e *Employee) GiveHike(percent float64) {
	e.Salary = e.Salary + (e.Salary * (percent / 100))
}

*/


// 2.28.44