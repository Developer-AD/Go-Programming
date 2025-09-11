package main

import (
	"fmt"
)

// // order struct
// type orderStruct struct {
// 	id        int
// 	amount    float32
// 	status    string
// 	isPaid    bool
// 	createdAt time.Time // Nano second precision.
// }

// // Mimic the constructor
// func newOrderStruct(id int, amount float32, status string) *orderStruct {
// 	order := orderStruct{
// 		id:     id,
// 		amount: amount,
// 		status: status,
// 	}
// 	return &order
// }

// // add * to make changes [receiver type.]
// func (o *orderStruct) changeStatus(status string) {
// 	o.status = status
// }

// // remove * if not changing [receiver type.]
// func (o orderStruct) getTotAmount() float32 {
// 	return o.amount + 100
// }

type customerStruct struct {
	name  string
	phone string
}

type orderStruct struct {
	id       int
	price    float32
	customer customerStruct
}

func main() {
	fmt.Println("------ Struct -----")

	// order := orderStruct{
	// 	id:     1,
	// 	amount: 199.50,
	// 	status: "created",
	// 	isPaid: true,
	// }

	// fmt.Println("before Order -", order)
	// order.createdAt = time.Now()l
	// order.changeStatus("delivered")
	// fmt.Println("after Order -", order)

	// fmt.Println("Order id -", order.id)
	// fmt.Println("Order amount -", order.amount)
	// fmt.Println("Order status -", order.status)
	// fmt.Println("Order isPaid -", order.isPaid)
	// fmt.Println("Order getTotAmount -", order.getTotAmount())
	// fmt.Println("Order createdAt -", order.createdAt)

	// order := *newOrderStruct(1, 199, "created")
	// fmt.Println("Constructor mimic -", order)

	// struct without name: inline struct for using single time.
	// languages := struct {
	// 	name   string
	// 	isGood bool
	// }{"Go-Lang", true}

	// fmt.Println("Languages :", languages)

	// Struct embedding. or composition
	customer := customerStruct{
		name:  "Abhishek",
		phone: "9399938409",
	}

	// order := orderStruct{
	// 	id:       1,
	// 	price:    199,
	// 	customer: customer,
	// }

	order := orderStruct{
		id:    1,
		price: 199,
		customer: customerStruct{
			name:  "Abhishek",
			phone: "9399938409",
		},
	}

	fmt.Println("Customer: ", customer)
	fmt.Println("Order: ", order)
	fmt.Println("cust: ", order.customer)
	// order.customer.name = "Abhishek"
	// order.customer.phone = "12345"
	fmt.Println("values added in cust: ", order.customer)
	fmt.Println("name: ", order.customer.name)
	fmt.Println("phone: ", order.customer.phone)

}
