package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to range.")

	// Iterate over slice using range.

	nums := []int{1, 2, 3, 4, 5}

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println("Nums values -", nums[i])
	// }

	sum := 0
	for ind, num := range nums {
		fmt.Printf("Index - %d & num - %d \n", ind, num)
		sum += num
	}

	fmt.Println("Total sum =", sum)

	// Iterate over map using range.
	data := map[string]string{"fname": "Abhishek", "lname": "Kumar"}

	// returns key, val
	for key, val := range data {
		fmt.Println("key-", key, " & val-", val)
	}

	// returns key
	for key := range data {
		fmt.Println("key-", key)
	}

	// iterate over string : rune
	// str_byte_ind: starting byte index of rune not a index.
	for str_byte_ind, unicode := range "Abhishek" {
		fmt.Println("str_byte_ind-", str_byte_ind, " unicode-", unicode, " &Char -", string(unicode))
	}

}
