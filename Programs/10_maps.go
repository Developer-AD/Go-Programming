package main

import (
	"fmt"
	"maps"
)

func main() {
	fmt.Println("------------ Map in go ---------")
	m := make(map[string]string)
	nums := make(map[string]int)

	nums["1"] = 1
	nums["2"] = 2

	fmt.Println("m - ", m)
	m["name"] = "Abhishek Kumar"
	m["address"] = "Hyderabad"
	m["age"] = "21"
	fmt.Println("m - ", m)
	fmt.Println("name - ", m["name"])
	fmt.Println("address - ", m["address"])
	fmt.Println("age - ", m["age"])
	fmt.Println("key - ", m["key"])
	fmt.Println("nums - ", nums)
	fmt.Println("nums-1 - ", nums["1"])
	fmt.Println("nums-2 - ", nums["2"])
	fmt.Println("nums-3 - ", nums["3"])
	fmt.Println("Len m - ", len(m))

	// delete(m, "address")
	// fmt.Println("After delete address - ", m)
	// clear(m)

	fmt.Println("After clear m - ", m)

	// mp := map[string]int{"age": 27}
	// mp := map[string]int{}
	// mp := map[string]int{}
	// fmt.Println("mp - ", mp)

	// val, ok := m["name"]
	val, key := m["name"]

	fmt.Println("Val -", val, ", key -", key)
	// if key {
	// 	fmt.Println("Key present.")
	// } else {
	// 	fmt.Println("Key not present")
	// }
	if !key {
		fmt.Println("Key Not present.")
	} else {
		fmt.Println("Key present")
	}

	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"a": 1, "b": 1}

	fmt.Println("Map-1", m1)
	fmt.Println("Map-2", m2)

	fmt.Println(maps.Equal(m1, m2))
}
