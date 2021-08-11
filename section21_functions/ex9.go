package main

import "fmt"

func printInt (x int) {
	fmt.Println(x)
}

func main() {
	fn := printInt
	fmt.Printf("%T\n", fn)
	fn(3)
}
