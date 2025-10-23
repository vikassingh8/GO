package main
import "fmt"

type nums[T any] struct {
	a []T
	
}




func num [T int | string](a, b T) (T, T) {
	return a, b

}


func main() {
	a, b := num(5, 10)
	fmt.Println("Int values:", a, b)
	
}