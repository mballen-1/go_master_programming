package main

import "fmt"

func main3() {
	type Person struct{
		name string
		age int8
		favoriteColors []string
	}
	me := Person{
		name : "Kike",
		age: 23,
		favoriteColors: []string{
			"Black", "White",
		},
	}

	you := Person{
		name : "Jack Sparrow",
		age: 32,
		favoriteColors: []string{
			"Red", "Blue",
		},
	}

	for i, c := range me.favoriteColors{
		fmt.Printf(" color at position %v is = %v\n", i, c)
	}

	_ = you
}
