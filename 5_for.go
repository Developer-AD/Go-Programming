package main

import "fmt"

func main() {
	fmt.Println("------------- For loop -------------")

	// In while loop fashion.
	// i := 1
	// for i <= 10 {
	// 	fmt.Println("i -", i)
	// 	// i = i + 1
	// 	// i += 1
	// 	i++
	// }

	// // Infinite loop.
	// i := 1
	// for true {
	// 	fmt.Println(i)
	// 	i++
	// }

	// // Infinite loop.
	// i := 1
	// for {
	// 	fmt.Println(i)
	// 	i++
	// }

	// // Normal for loop.
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println("i -", i)
	// }

	// // Continue and Break.
	// for i := 1; i <= 10; i++ {
	// 	// if i == 5 {
	// 	// 	continue
	// 	// }
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println("i -", i)
	// }

	// range with for
	for i := range 5 { // range(5) = 0, 1, 2, 3, 4, 5
		fmt.Println("for i -", i)
	}

	fmt.Scanln() // Holds the exe for key input.

}
