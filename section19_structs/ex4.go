package main

import "fmt"

func main() {

	type Grades struct{
		grade, course string
	}

	type Person struct{
		name string
		age int8
		favoriteColors []string
		grades Grades
	}

	me := Person{
		name : "Kike",
		age: 23,
		favoriteColors: []string{
			"Black", "White",
		},
		grades: Grades{
			grade: "A+",
			course: "data structures",
		},
	}

	you := Person{
		name : "Jack Sparrow",
		age: 32,
		favoriteColors: []string{
			"Red", "Blue",
		},
		grades: Grades{
			grade: "A+",
			course: "Algorithms",
		},
	}

	you.grades.grade = "Golang"
	you.grades.course = "98"

	fmt.Printf("me: \n%#v\n", me)
	fmt.Printf("you: \n%#v\n", you)


}
