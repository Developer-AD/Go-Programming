package main

import "fmt"

// Send value only.
// func changeNum(num int) {
// 	num = 100
// 	fmt.Println("New number inside fun - ", num)
// 	fmt.Println("Mem Addr inside a func-", &num)
// }

func changeNumPointer(num *int) {
	*num = 100
	fmt.Println("New number inside fun - ", *num)
	fmt.Println("Mem Addr inside a func-", &num)
}

func main() {
	fmt.Println("--------- Pointers ---------")

	num := 1

	fmt.Println("Num-", num)
	fmt.Println("Mem Addr-", &num)
	fmt.Println("Num in main fun before fun call.")
	// changeNum(num)
	changeNumPointer(&num)
	fmt.Println("Num in main fun After fun call.")
	fmt.Println("Num-", num)
}
