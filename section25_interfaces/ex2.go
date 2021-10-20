package main

import "fmt"

func main2() {
	var empty interface{}
	empty = 234
	fmt.Printf("empty is = %#v\n", empty)

	empty = 333.4
	fmt.Printf("empty is = %#v\n", empty)

	empty = []int{3, 5, 5, 5}
	fmt.Printf("empty is = %#v\n", empty)

	e, ok := empty.([]int)
	if ok == true {
		empty = append(e, 2)
	}
	fmt.Printf("empty is = %#v\n", empty)
}
