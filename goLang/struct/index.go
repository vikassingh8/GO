package main
import "fmt"

type Order struct {
	id      int
	amount  int
	status   string
}


func (o *Order) statusMethod(status string){
	o.status=status
}

func newOrder (id int,amount int,status string)*Order{
	order:=Order{id,amount,status}
	return & order
}
func main() {
	result:=newOrder(1,200,"vikas")
	fmt.Println(*result)
	// order := Order{id: 1}
	order2 :=Order{id: 2, amount: 100, status: "pending"}
	order2.statusMethod("vikas")
	


	fmt.Println("Order 2:", order2.status)
	// fmt.Println("Order ID:", order.id)


}
