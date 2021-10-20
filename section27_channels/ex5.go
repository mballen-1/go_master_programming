package main

import (
	"fmt"
)

func main() {
	ch := make(chan float64)
	for i := 1.; i <= 100; i++ {
		go func(x float64) {
			ch <- x * x
		}(i)
	}

	for i := 1; i <= 100; i++ {
		fmt.Println(<-ch)
	}
}
