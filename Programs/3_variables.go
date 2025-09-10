package main

import "fmt"

/*
"""
--------------- Data Type in Go ---------------
int
float
string
bool
"""
*/
// var code = "IN+91" // Correct
// code := "IN+91" // Error: non-declaration statement outside function body

func main() {
	var name string = "Abhishek Kumar"
	// var age = 27
	age := 27
	var balance = 100.50
	fmt.Printf("Hello, %s", name)
	// fmt.Printf("Hello, %d", name) // Incorrect.
	fmt.Printf("\nYour age is - %d", age)
	fmt.Printf("\nAc bal - %g\n", balance)

	var company string // When we define the var and use later.
	company = "Google"

	// Shorthand - declaring and initializing a variable
	new_comp := "Microsoft"

	fmt.Println("Company - ", company)
	fmt.Println("New Company - ", new_comp)

	// You can declare multiple variables at once.
	// var a, b, c int
	// a = 10
	// b = 20
	// c = 30

	var a, b, c int = 10, 20, 30

	sum := a + b + c
	fmt.Println("Sum of a, b and c is =", sum)
	fmt.Printf("Sum of a, b and c is = %d", sum)

	code := "IN+91"
	fmt.Println("\nCountry code -", code)
}
