// package main
// import (
// 	"context"
// 	"fmt"
// 	"time"
// )


// func worker(ctx context.Context) {
// 	for {
// 		select {

// 		case <-ctx.Done():
// 			fmt.Println("Worker stopped")
// 			return

// 		default:
// 			fmt.Println("Working...")
// 			time.Sleep(time.Second)
// 		}
// 	}
// }
// func main(){
// 	ctx, cancel := context.WithCancel(context.Background())

// go worker(ctx)


// defer cancel()
// }

package main

import (
	"context"
	"fmt"
	"time"
)

func fetchData(ctx context.Context) {

	select {

	case <-time.After(5 * time.Second):
		fmt.Println("Data fetched")

	case <-ctx.Done():
		fmt.Println("Request cancelled:", ctx.Err())
	}
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go fetchData(ctx)

	time.Sleep(6 * time.Second)
}