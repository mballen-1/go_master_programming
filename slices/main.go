package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	lenArgs := len(os.Args)
	if lenArgs >=3 && lenArgs <= 11 {
		sum, prod := 0., 1.
		for _, a := range os.Args {
			if string(a[0]) == "C" {
				continue
			}
			fmt.Println("a = " , a)
			p, _ := strconv.ParseFloat(a, 64)
			sum += p
			prod  *= p
		}
		fmt.Printf("Sum = %.2f, Prod = %.2f \n", sum, prod)
	} else {
		fmt.Printf("Args not in range")
	}
}
