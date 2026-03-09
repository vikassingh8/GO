// package main

// import "fmt"

// type nums[T any] struct {
// 	a []T
// }
//  func num[T int | string](a, b T) (T, T) {
// 	return a, b
// }

// func main() {
// 	a, b := num(5, 10)
// 	fmt.Println("Int values:", a, b)
// }



package main

import "fmt"

// Generic struct
type nums[T any] struct {
	a []T
}
// Generic function
func num[T int | string](a, b T) (T, T) {
	return a, b
}

func main() {
	// Using generic struct with int
	n := nums[int]{[]int{1, 2, 3}}
	fmt.Println("Nums struct:", n)

	// Using generic function
	x, y := num(5, 10)
	fmt.Println("Int values:", x, y)

	// Using generic struct with string
	s := nums[string]{a: []string{"go", "lang"}}
	fmt.Println("String struct:", s.a)
}
