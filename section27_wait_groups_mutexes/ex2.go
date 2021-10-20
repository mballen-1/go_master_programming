package main

import (
	"fmt"
	"sync"
)

func sum(a, b float64, wg *sync.WaitGroup){
	fmt.Printf("sum is = %.2f\n", a + b)
	wg.Done()
}

func main2() {
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i< 3; i++{
		go sum(3.2, float64(i * 10), &wg)
	}
	wg.Wait()
}
