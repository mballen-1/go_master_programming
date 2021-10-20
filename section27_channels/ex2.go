package main

import "fmt"

func main2() {
	ch := make(chan string)
	go func(v string, c *chan string) {
		*c <- v
	}("this is the value", &ch)
	fmt.Println(<-ch)
}
