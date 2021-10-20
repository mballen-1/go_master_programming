package main

import (
	"fmt"
	"sync"
)

func deposit(b *int, n int, wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
		*b += n
		wg.Done()
	mu.Unlock()
}

func withdraw(b *int, n int, wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
		*b -= n
		wg.Done()
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(200)

	balance := 100

	for i := 0; i < 100; i++ {
		go deposit(&balance, i, &wg, &mu)
		go withdraw(&balance, i, &wg, &mu)
	}
	wg.Wait()

	fmt.Println("Final balance value:", balance)
}

func mySync(){

}