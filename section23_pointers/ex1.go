package main

import "fmt"

func main1() {
	// fmt.Println("Hello, playground")
	x := 10.10
	fmt.Printf("The adress of x is %p\n", &x)

	ptr := &x
	fmt.Printf("type of ptr is %T, and its value is %f\n", ptr, *ptr)

	fmt.Printf("Address of the pointer ptr is %v, The value of x through pointer is %f\n", &ptr, *ptr)
}
