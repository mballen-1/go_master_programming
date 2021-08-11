package main

import (
	"fmt"
)

func sum(vls ...int) (res int){
	for _, v := range vls{
		res += v
	}
	return
}


func main4() {
	fmt.Println(sum(3,4,5,6,7))
}
