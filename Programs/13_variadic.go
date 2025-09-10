package main

import "fmt"

// n numbers of parameter in int
func sum(nums ...int) int {
	s := 0
	for _, num := range nums {
		fmt.Println("num-", num)
		s += num
	}
	return s
}

// n numbers of parameter in int
func any_params(items ...interface{}) {
	for ind, item := range items {
		fmt.Println(ind, "--->", item)
	}

}

func main() {
	fmt.Println("----- Variadic function -----")

	// fmt.Println(1, 3, 4, "name") // This is also a veriadic function.
	// sums := sum(1, 2, 3, 4, 5)
	numbers := []int{1, 2, 1, 2, 4}
	sums := sum(numbers...)
	fmt.Println(sums)

	fmt.Println("Any---params.")
	any_params(1, 2, "Abhishek", 100.55, true)
}
