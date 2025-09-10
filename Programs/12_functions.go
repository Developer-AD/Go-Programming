package main

import "fmt"

// func add(x int, y int) int {
// 	return x + y
// }

func add1(x, y int) int {
	return x + y
}

func greet(name string) string {
	return "Hello " + name
}

func getNames() (string, string, string) {
	return "Abhishek", "Pradeep", "Nitesh"
}

// Advanced Function concepts.
func processIt(fn func(a int) int) {
	fn(100)
}

func retFunc() func(a int) int {
	return func(a int) int {
		return 98
	}
}

func main() {
	fmt.Println("-------------- Functions --------------")

	// addition := add(2, 5)
	addition := add1(3, 5)
	fmt.Println("Addition is: ", addition)
	fmt.Println("Greeting : ", greet("Abhishek"))

	// Multiple returns
	// name1, name2, name3 := getNames()
	name1, _, _ := getNames()
	// fmt.Println("Name1-", name1, ", Name2-", name2, ", Name3-")
	fmt.Println("Name1-", name1, ", Name2-", ", Name3-")

	// Function.
	// fn := func(a int) int {
	// 	return 99
	// }

	// processIt(fn)
	// fmt.Println("Process -")

	ret := retFunc()
	fmt.Println("Return fn -", ret)
}
