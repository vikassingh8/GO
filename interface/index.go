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

// package main

// import "fmt"

// // Define an interface (behavior)
// type Creator interface {
// 	Create(name string, age int)
// }

// // Struct
// type Student struct{}

// // Method implementation
// func (s Student) Create(name string, age int) {
// 	fmt.Println("Student Name:", name, "Age:", age)
// }

// // Function that depends on interface, not concrete type
// func createEntity(c Creator, name string, age int) {
// 	c.Create(name, age)
// }

// func main() {
// 	var c Creator
// 	c = Student{}          // Student implicitly implements Creator
// 	createEntity(c, "Alice", 20)
// }

package main

import "fmt"

type UserRepository interface {
	CreateUser(email string)
	GetUser(email string)
}

type userRepo struct{}

func (u userRepo) CreateUser(email string) {
	fmt.Println("Repository: inserting user into database:", email)
}

func (u userRepo) GetUser(email string) {
	fmt.Println("Repository: fetching user from database:", email)
}
// services

type AuthService interface {
	Register(email string)
	Login(email string)
}

type authService struct {
	repo UserRepository
}

func (a authService) Register(email string) {
	fmt.Println("Service: validating and preparing user")
	a.repo.CreateUser(email)
}

func (a authService) Login(email string) {
	fmt.Println("Service: checking login")
	a.repo.GetUser(email)
}

// hander

type AuthHandler struct {
	service AuthService
}

func (h AuthHandler) RegisterHandler(email string) {
	fmt.Println("Handler: received register request")
	h.service.Register(email)
}

func (h AuthHandler) LoginHandler(email string) {
	fmt.Println("Handler: received login request")
	h.service.Login(email)
}


func main() {

	repo := userRepo{}

	service := authService{
		repo: repo,
	}

	handler := AuthHandler{
		service: service,
	}

	handler.RegisterHandler("john@example.com")

	handler.LoginHandler("john@example.com")
}