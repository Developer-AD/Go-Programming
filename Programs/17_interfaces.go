package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	// gateway stripe
	// gateway razorpay
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	// rp := razorpay{}
	// rp.pay(amount)

	// st := stripe{}
	// st.pay(amount)

	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("Razorpay payment of", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("Stripe payment of", amount)
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("Paypal payment of", amount)
}

func main() {
	// st := stripe{}
	// rp := razorpay{}
	p := paypal{}
	// newPayment := payment{gateway: st}
	// newPayment := payment{gateway: rp}
	newPayment := payment{gateway: p}
	// newPayment := payment{gateway: st}
	newPayment.makePayment(199)
}
