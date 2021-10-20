package main

import "fmt"

func main1() {
	var c1 chan float64
	c2 := make(<-chan rune) // recieve only
	c3 := make(chan<- rune) // send only
	c4 := make(chan int, 10)

	fmt.Printf("c1 type is %T\n", c1)
	fmt.Printf("c2 type is %T\n", c2)
	fmt.Printf("c3 type is %T\n", c3)
	fmt.Printf("c4 type is %T\n", c4)
}
