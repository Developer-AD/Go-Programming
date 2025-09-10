package main

import "fmt"

func main() {
	fmt.Println(1) // Integer
	fmt.Println(1 + 100)
	fmt.Println('A')                        // Single quote only single char (Literal)
	fmt.Println("Abhishek Kumar")           // String
	fmt.Println("Abhishek" + " " + "Kumar") // String
	fmt.Println(true)                       // bool values
	fmt.Println(false)                      // bool values
	fmt.Println(10.55)                      // float values
	fmt.Println(99.123456789)               // float values
	fmt.Println(2 / 3)                      // Operation on int values -> 0
	fmt.Println(2 / 3.0)                    // Operation on float valuesv -> 0.6666666666666666

	// Logical Operations.
	fmt.Println("-----------------------------")
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(!false)
}
