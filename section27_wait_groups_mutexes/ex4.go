package main

import (
	"fmt"
	"math"
	"sync"
)

func main4() {
	var wg sync.WaitGroup
	wg.Add(50)
	for i := 100.; i < 150; i++{
		go func (x float64, wg *sync.WaitGroup) {
			fmt.Printf("sqrt of %.0f is %.4f\n", x,math.Sqrt(x))
			wg.Done()
		}(i, &wg)
	}
	wg.Wait()
}