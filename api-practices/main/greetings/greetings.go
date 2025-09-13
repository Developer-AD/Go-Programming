package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// func RandomGreetingFormats() string { // can use outside this packages as well

// it's not exported: only available withing this package.
func randomGreetingFormats() string {
	formats := []string{
		"Welcome %v.",
		"Hello %v, good to see you here.",
		"Good morning %v, have a nice day",
	}
	// fmt.Println("rand-int", rand.Int())
	// fmt.Println("rand-intn-10", rand.Intn(10))
	return formats[rand.Intn(len(formats))]
}

func Greet(name string) (string, error) {
	if name == "" {
		return name, errors.New("name is empty")
	}
	// message := fmt.Sprintf("Hello %v, Welcome to go programming.", name)
	format := randomGreetingFormats()
	message := fmt.Sprintf(format, name)
	return message, nil
}
