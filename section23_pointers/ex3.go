package main

import "fmt"

func swapByPointers(a *float64, b *float64) (float64, float64){
	return *b, *a
}

func main() {
	x, y := 5.5, 8.8
	x, y = swapByPointers(&x, &y)
	fmt.Printf("x = %f, y = %f \n", x, y)
}
