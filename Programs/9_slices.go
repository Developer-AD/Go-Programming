package main

import "fmt"

func main() {
	// Slices
	var names []string
	fmt.Println(names)
	fmt.Println(names == nil)
	fmt.Println(len(names))

	// var nums = make([]int, 3)
	// var nums = make([]int, 2, 10)
	var nums = make([]int, 0, 10) // Empty slice.
	fmt.Println(nums == nil)
	fmt.Println(nums)
	fmt.Println(len(nums))

	nums = append(nums, 100)
	nums = append(nums, 200)

	fmt.Println("Capacity -", cap(nums))
	fmt.Println("Elements -", nums)
	fmt.Println("Size -", len(nums))

	elements := []string{} // not nil
	// var elements []string // nil

	fmt.Println("Elements - ", elements)
	fmt.Println("nil", elements == nil)
	fmt.Println("size", len(elements))

	cp1 := []int{}
	cp2 := []int{}

	cp1 = append(cp1, 10)
	cp1 = append(cp1, 20)
	copy(cp2, cp1)

	fmt.Println("cp1 -", cp1)
	fmt.Println("cp2 -", cp2)
}
