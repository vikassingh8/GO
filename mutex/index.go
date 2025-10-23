package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) increment(wg *sync.WaitGroup) {
	defer wg.Done()     // Signal that this goroutine is done
	p.mu.Lock()          // Lock before updating shared data
	p.views += 1
	p.mu.Unlock()        // Unlock after done
}

func main() {
	var wg sync.WaitGroup
	data := post{views: 100}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go data.increment(&wg)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Total views:", data.views)
}
