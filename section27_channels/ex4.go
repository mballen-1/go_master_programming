package main

import (
	"fmt"
	"math"
)

func power(x float64, ch chan float64) {
	ch <- math.Pow(x, 2)
}

func main4() {
	ch := make(chan float64)
	for i := 1.; i <= 50; i++ {
		go power(i, ch)
	}

	for i := 1; i <= 50; i++ {
		fmt.Println(<-ch)
	}
}
