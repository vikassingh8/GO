package main
import "fmt"

type OrderStatus int

const (
	Pending OrderStatus = iota
	Processing
	Completed
	Cancelled
)

func main() {
	var status OrderStatus = Processing
	fmt.Println("Current Order Status:", status)
}