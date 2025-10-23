package main

import (
	"fmt"
	"sync"
)

// // func task (id int) {
// // 	fmt.Println( id)
// // }

// func main() {
// 	for i := 1; i <= 15; i++ {
// 		// go task(i)
// 		go func(i int) {
// 			fmt.Println(i)
// 		}(i)
// 	}
// 	time.Sleep(time.Second * 2)
// }

func task (id int,w *sync.WaitGroup){

	defer w.Done()
	fmt.Println("Task ID:", id)
}
func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 15; i++ {
		wg.Add(1)
		go task(i, &wg)
	}
	wg.Wait()
}