package main

import (
	"fmt"
)

func searchItem(slc []string, str string )bool{
	for _, item := range slc {
		if item == str {
			return true
		}
	}
	return false
}

func main6() {
	fmt.Println(searchItem([]string{"cuerpo", "desierto", "esfinge"}, "puerco"))
	fmt.Println(searchItem([]string{"cuerpo", "desierto", "esfinge"}, "cuerpo"))
}
