package main

import (
	"fmt"
)

func main() {
	fmt.Println("----------------- Switch -----------------")

	// num := 7

	// switch num {
	// // switch num {
	// case 0:
	// 	println("It's Sunday.")
	// case 1:
	// 	println("It's Monday.")
	// case 2:
	// 	println("It's Tuesday.")
	// case 3:
	// 	println("It's Wednesday.")
	// case 4:
	// 	println("It's Thursday.")
	// case 5:
	// 	println("It's Friday.")
	// case 6:
	// 	println("It's Saturday.")
	// default:
	// 	fmt.Println("Not a valid weekday.")
	// }

	// complex := time.Now().Weekday()

	// switch complex {
	// // switch num {
	// case time.Sunday, time.Saturday:
	// 	println("It's weekday.")
	// case time.Monday:
	// 	println("It's Monday.")
	// case time.Tuesday:
	// 	println("It's Tuesday.")
	// case time.Wednesday:
	// 	println("It's Wednesday.")
	// case time.Thursday:
	// 	println("It's Thursday.")
	// case time.Friday:
	// 	println("It's Friday.")

	// default:
	// 	fmt.Println("Not a valid weekday.")
	// }

	// t := time.Now()

	// switch {
	// case t.Hour() < 12:
	// 	{
	// 		fmt.Println("Before 12 PM in the afternoon.")
	// 	}

	// default:
	// 	fmt.Println("After 12 PM in the afternoon.")
	// }

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whoAmI(true)
	whoAmI(1)
	whoAmI("hey")

}
