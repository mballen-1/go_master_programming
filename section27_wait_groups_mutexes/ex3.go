package main

import (
	"fmt"
	"math"
	"sync"
)

func main3() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func (x float64, wg *sync.WaitGroup) {
		fmt.Printf("sqrt of %f is %f", x,math.Sqrt(x))
		wg.Done()
	}(3.0423, &wg)
	wg.Wait()
}
