package main

import "fmt"

type money float64

func (m money) print() {
	fmt.Printf("Money value is %.2f", m)
}

func main1() {
	var m money = 3.222
	m.print()
}
