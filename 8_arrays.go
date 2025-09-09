package main

import "fmt"

func main() {
	fmt.Println("--------------- Arrays Tutorial ----------")

	// Array default stores zeroed values. int-0 bool-false

	var nums [5]int
	var vals [5]bool
	var names [5]string

	numbers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Length of array: ", len(nums))
	fmt.Println("Items of nums: ", nums)
	fmt.Println("Items of vals: ", vals)
	fmt.Println("Items of names: ", names)

	nums[0] = 100
	// fmt.Println(nums[0])
	fmt.Println("Numbers -", numbers)

	// 2D Arrays.
	arrs := [2][2]int{{1, 2}, {3, 4}}

	fmt.Println(arrs)

	// Count the size - [...] on compile time.

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", len(b), b)

	c := [...]int{1, 2, 3, 4}
	fmt.Println("dcl:", len(c), c)

	d := [3]int{1, 2, 3}
	fmt.Println("dcl:", len(d), d)
}
