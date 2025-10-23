// package main

// import "fmt"

// // Step 1: Define a common interface
// type paymenter interface {
// 	pay(amt float64)
// }

// // Step 2: Define Stripe gateway
// type stripe struct{}

// func (s stripe) pay(amt float64) {
// 	fmt.Println("💳 Stripe: Paid", amt)
// }

// // Step 3: Define Razorpay gateway
// type razorpay struct{}

// func (r razorpay) pay(amt float64) {
// 	fmt.Println("💰 Razorpay: Paid", amt)
// }

// // Step 4: Payment struct that can use any gateway
// type payment struct {
// 	gateway paymenter
// }

// // Step 5: Process payment through selected gateway
// func (p payment) processPayment(amt float64) {
// 	p.gateway.pay(amt)
// }

// // Step 6: Main program
// func main() {
// 	// Use Stripe
// 	stripeGateway := stripe{}
// 	p1 := payment{gateway: stripeGateway}
// 	p1.processPayment(100.0)

// 	// Use Razorpay
// 	razorpayGateway := razorpay{}
// 	p2 := payment{gateway: razorpayGateway}
// 	p2.processPayment(250.0)
// }





package main
import "fmt"


// Step 1: Define a common interface
type method interface {
	craete(name string, age int)
}

type student struct{}

func (s student) craete(name string, age int) {
	fmt.Println("Student Name:", name, "Age:", age)
}



func main() {
	student1 := student{}
	student1.craete("Alice", 20)


}