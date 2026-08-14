
# GoLang Learning Journey

[![Go Version](https://img.shields.io/badge/Go-1.26.2-00ADD8?style=flat-square&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg?style=flat-square)](LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-Welcome-brightgreen.svg?style=flat-square)](CONTRIBUTING.md)
[![Platform](https://img.shields.io/badge/Platform-Cross--Platform-blue?style=flat-square)](#)

Welcome to the **GoLang Learning Journey** repository! This project is a structured, hands-on curriculum designed to take developers from absolute beginners to highly proficient Go (Golang) programmers. 

Each directory represents a self-contained, runnable module focusing on a core Go concept, complete with practical code examples, idiomatic patterns, and detailed explanations.

---

## 🗺️ Learning Roadmap & Directory Structure

The repository is organized sequentially to build your knowledge step-by-step. Below is the breakdown of each module and the concepts covered:

| Module | Directory | Key Concepts Covered |
| :--- | :--- | :--- |
| **01** | [`01_setup_first_program/`](./01_setup_first_program/) | Go workspace setup, `package main`, entry point `main()`, and basic console output. |
| **02** | [`02_variables_and_types/`](./02_variables_and_types/) | Static typing, explicit variable declarations, and default zero-values. |
| **03** | [`03_packages_imports/`](./03_packages_imports/) | Importing standard library packages, package-level scope, and formatting with `fmt`. |
| **04** | [`04_var_vs_short_declare/`](./04_var_vs_short_declare/) | Comparing `var` vs. the short declaration operator (`:=`), block scope vs. package scope. |
| **05** | [`05_basic_types_string/`](./05_basic_types_string/) | String manipulation, immutability, raw string literals, and basic string operations. |
| **06** | [`06_basic_types_int_float/`](./06_basic_types_int_float/) | Numeric types (`int`, `int8` to `int64`, `float32`, `float64`), precision, and type casting. |
| **07** | [`07_basic_types_boolean/`](./07_basic_types_boolean/) | Boolean logic, comparison operators, and conditional flags. |
| **08** | [`08_constants/`](./08_constants/) | Typed and untyped constants, compile-time evaluation, and the `const` keyword. |
| **09** | [`09_if_else/`](./09_if_else/) | Conditional branching, logical operators, and short-statement initialization in `if` blocks. |
| **10** | [`10_for_classic_loop/`](./10_for_classic_loop/) | Go's single looping construct: classic three-component `for` loops, and `while`-equivalent loops. |
| **11** | [`11_switch_statement/`](./11_switch_statement/) | Expression switches, tagless switches, multi-case evaluation, and implicit break behavior. |
| **12** | [`12_arrays/`](./12_arrays/) | Fixed-size sequences, memory allocation, array initialization, and length limitations. |
| **13** | [`13_slice_literal/`](./13_slice_literal/) | Dynamic arrays, slice literals, and basic slicing operations. |
| **14** | [`14_slice_length_and_capacity/`](./14_slice_length_and_capacity/) | Internal mechanics of slices: length (`len`), capacity (`cap`), and backing arrays. |
| **15** | [`15_for_range_over_slice/`](./15_for_range_over_slice/) | Idiomatic iteration over slices using the `range` keyword, index, and value extraction. |
| **16** | [`16_map/`](./16_map/) | Hash maps, key-value pairs, dynamic allocation with `make()`, element deletion, and the **comma-ok** idiom. |
| **17** | [`17_functions/`](./17_functions/) | Function declarations, multiple return values, named (naked) returns, and variadic functions (`...int`). |
| **18** | [`18_anonymous_function/`](./18_anonymous_function/) | Anonymous functions, Immediately Invoked Function Expressions (IIFE), inline argument passing, and assigning functions to variables. |
| **19** | [`19_error_return_pattern/`](./19_error_return_pattern/) | Idiomatic error handling, returning errors as normal values, custom error formatting with `fmt.Errorf`, and error propagation. |
| **20** | [`20_defer_basics/`](./20_defer_basics/) | The `defer` statement, scheduling function execution, and guaranteeing resource cleanup across all return paths. |
| **21** | [`21_pointers/`](./21_pointers/) | Introduction to pointers, memory addresses, and dereferencing in Go. |
| **22** | [`22_structs/`](./22_structs/) | Struct definition, named vs. positional initialization, zero values, struct pointers, embedding (composition), struct tags (JSON), and method receivers. |
| **23** | [`23_methods_val_receiver/`](./23_methods_val_receiver/) | Methods with value receivers, receiving copies of structs, and read-only operations. |
| **24** | [`24_methods_pointer_receiver/`](./24_methods_pointer_receiver/) | Methods with pointer receivers, modifying struct fields, and avoiding memory copies. |
| **25** | [`25_golang_module/`](./25_golang_module/) | Go modules, dependency management, project layout (`cmd/` and `internal/`), and package visibility. |
| **26** | [`26_golang_net_http_module/`](./26_golang_net_http_module/) | Standard library `net/http` module, HTTP servers, routing, JSON encoding/decoding, HTTP GET requests, JSON unmarshalling, and consuming external APIs. |
| **27** | [`27_practise/`](./27_practise/) | Comprehensive practice module consolidating Go fundamentals: variables, basic types, constants, control structures, arrays, slices, maps, and pointers. |

---
## 🛠️ Tech Stack & Requirements

*   **Language:** Go (Golang)
*   **Minimum Go Version:** `1.26.2` (as specified in `go.mod`)
*   **Operating System:** Cross-platform (Linux, macOS, Windows)
*   **IDE Recommendation:** VS Code (with the official Go extension) or GoLand
*   **Dependencies:** Standard library focused (external dependencies like JWT and bcrypt have been removed from the Todo REST API to leverage native Go capabilities)

---
## 🚀 Getting Started

### Prerequisites

Ensure you have Go installed on your system. You can verify your installation by running:

bash
go version


If Go is not installed, download it from the [official Go downloads page](https://go.dev/dl/).

### Installation

1. Clone this repository to your local machine:
   bash
   git clone https://github.com/Anos714/GoLang_Learning.git
   

2. Navigate into the project directory:
   bash
   cd GoLang_Learning
   

3. Initialize/verify the module dependencies:
   bash
   go mod tidy
   
## 📖 Usage & Execution

Every module in this repository is designed to be executed independently. To run any specific module, navigate to its directory and run the `main.go` file (or the main entry point).

### Example 1: Running a Basic Module (e.g., Maps)

To run the `16_map` module:

bash
cd 16_map
go run main.go


### Example 2: Running the Todo REST API Module

The `28_todo_rest_api_using_nethttp` module implements a REST API. To run it:

1. Navigate to the module's directory:
   bash
   cd 28_todo_rest_api_using_nethttp
   

2. Create and configure a `.env` file in this directory (or the server will fall back to system environment variables and default port `8000`).
3. Run the application:
   bash
   go run main.go
   

4. You can test the server by sending a GET request to the ping endpoint:
   bash
   curl http://localhost:8000/ping
   

### Example 3: Running the Gin REST API Module (with MongoDB)

The `29_gin_rest_api` module implements a Notes API using the Gin framework and MongoDB. To run it:

1. Navigate to the module's directory:
   bash
   cd 29_gin_rest_api
   

2. Ensure you have MongoDB running and configure your environment variables (such as MongoDB URI and Database name).
3. Run the application:
   bash
   go run cmd/api/main.go
   
   Or, if you have [Air](https://github.com/air-verse/air) installed for live reloading:
   bash
   air
   
# Navigate to the module directory
cd 16_map

# Execute the program
go run main.go
```

### Code Spotlight: Idiomatic Go Maps (`16_map/main.go`)

Here is a preview of the concepts covered in our advanced modules, showcasing map initialization, safe key lookups using the **comma-ok** pattern, and iterations:

```go
package main

import "fmt"

func main() {
	// 1. Map Initialization with Literals
	votersAge := map[string]int{
		"rahul":  21,
		"hemant": 24,
		"ballu":  52,
		"kallu":  18,
	}
	fmt.Println("Voters:", votersAge)

	// 2. Dynamic Allocation using make()
	studentGrades := make(map[string]int)
	studentGrades["Rahul"] = 20
	studentGrades["Hemant"] = 22

	// 3. Deleting Keys safely
	delete(studentGrades, "Rahul")

	// 4. The Comma-Ok Idiom (Safe Lookups)
	if grade, ok := studentGrades["Hemant"]; ok {
		fmt.Printf("Found Hemant! Grade: %d\n", grade)
	} else {
		fmt.Println("Record not found")
	}

	// 5. Iterating over Maps
	total := 0
	for student, grade := range studentGrades {
		fmt.Printf("Student: %s | Grade: %d\n", student, grade)
		total += grade
	}
	fmt.Println("Total Class Grade:", total)
}
```

---

## 🧪 Running Tests

While these are learning modules, you can easily write and run tests. To run tests across the entire workspace (if added in future modules):

```bash
go test ./... -v
```

---

## 🤝 Contributing

Contributions are welcome! If you want to add new modules, fix bugs, or improve documentation:

1. **Fork** the repository.
2. Create a new branch for your feature:
   ```bash
   git checkout -b feature/amazing-new-module
   ```
3. Commit your changes with clear, descriptive messages:
   ```bash
   git commit -m "feat: add module 17 covering struct pointers"
   ```
4. Push your branch:
   ```bash
   git push origin feature/amazing-new-module
   ```
5. Open a **Pull Request** explaining your changes.

---

## 📝 License

This project is licensed under the **MIT License**. Feel free to use, modify, and distribute this codebase for educational and commercial purposes. See the [LICENSE](LICENSE) file for details.

---

## 🌟 Acknowledgments

*   The Go team for creating a simple, fast, and highly concurrent programming language.
*   All Go community members who contribute to idiomatic patterns and documentation.

---
*Happy Coding! Keep learning Go!* 🚀