package main

import (
	"fmt"
	"strings"
)

func searchItemInsensitive(slc []string, str string ) bool {
	for _, item := range slc {
		if strings.EqualFold(item, str) {
			return true
		}
	}
	return false
}

func main7() {
	fmt.Println(searchItemInsensitive([]string{"cuerpo", "desierto", "esfinge"}, "puerco"))
	fmt.Println(searchItemInsensitive([]string{"cuerpo", "desierto", "esfinge"}, "CueRpo"))
}
