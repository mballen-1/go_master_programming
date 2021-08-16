package main

import "fmt"

func main2() {
	x, y := 10, 5
	ptrx, ptry := &x, &y
	if y != 0 {
		z := *ptrx / *ptry
		fmt.Printf("z is %d", z)
	}
}
