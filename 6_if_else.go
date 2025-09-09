package main

import (
	"fmt"
)

func main() {
	println("----------- If else ----------------")
	// age := 16

	// if age >= 18 {
	// 	fmt.Println("Eligible for voting.")
	// } else {
	// 	fmt.Println("Not Eligible for voting.")
	// }

	// score := 30

	// if score >= 90 {
	// 	fmt.Println("Grade A++")

	if score := 98; score >= 90 {
		fmt.Println("Grade A++")
	} else if score >= 80 {
		fmt.Println("Grade A+")
	} else if score >= 60 {
		fmt.Println("Grade A")
	} else if score >= 50 {
		fmt.Println("Grade B")
	} else if score >= 33 {
		fmt.Println("Grade C")
	} else {
		println("Fail")
	}

	age := 60

	if age > 0 && age < 18 {
		fmt.Println("Teenager")
	} else if age >= 18 && age < 60 {
		fmt.Println("Adult")
	} else if age >= 60 {
		fmt.Println("Old.")
	}
}
