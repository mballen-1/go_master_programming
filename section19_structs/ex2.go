package main

import "fmt"

func main2() {
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
	me.name = "Andrei"
	you.favoriteColors = append(you.favoriteColors, "White")

	fmt.Printf("me: \n%#v\n", me)
	fmt.Printf("you: \n%#v\n", you)
}
