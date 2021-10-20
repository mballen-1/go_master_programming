package main

import "fmt"

type money float64

func (m money) print() {
	fmt.Printf("Money value is %.2f\n", m)
}

func (m money) printStr() string {
	return fmt.Sprintf("%.2f\n", m)
}

func main2() {
	var m money = 3.222
	m.print()
	fmt.Println(m.printStr())
}
