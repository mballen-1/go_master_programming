package main

import (
	"fmt"
	"strconv"
)

func myFunc(s string) int{
	fmt.Printf("s = %v\n", s)
	if len(s) > 0 {
		n, err := strconv.Atoi(s)
		if err != nil {
			p := fmt.Errorf("error with n conversion: %v", err)
			fmt.Print(p)
			return -1
		}
		nn, _ := strconv.Atoi(s + s)
		nnn, _ := strconv.Atoi(s + s + s)
		fmt.Printf("n = %v, nn = %v, nnn = %v\n", n, nn, nnn)
		return n + nn + nnn
	}
	p := fmt.Errorf("string is shorter than expected\n")
	fmt.Print(p)
	return -1
}


func main() {
	fmt.Println(myFunc("5"))
}
