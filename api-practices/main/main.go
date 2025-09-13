package main

import (
	"fmt"
	"log"

	"github.com/developer-ad/api-practices/greetings"
	// "github.com/developer-ad/api-practices/utils"
)

func main() {
	fmt.Println("-----------Welcome to main program -----------")

	// res := greetings.Greet("Abhishek Kumar")

	message, err := greetings.Greet("")
	// message, err := greetings.Greet("Abhishek Kumar")

	// check for errors.
	if err != nil {
		log.Fatal(err) // prints the error and exit the program.
		// fmt.Println(err) // only print the message
	}

	fmt.Println("Message -", message)
	fmt.Println("Error -", err)

	// if err == nil {
	// 	fmt.Println("----------- No errors ----------")
	// } else {
	// 	fmt.Println("----------- Some errors ----------")
	// }

	// ut := utils.GetData()
	// fmt.Println("utils:", ut)

	// fmt.Println("----------- Test Exports ----------")
	// msg := fmt.Sprintf(greetings.RandomGreetingFormats(), "Abhi")
	// fmt.Println(msg)
}
