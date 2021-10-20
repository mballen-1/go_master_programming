package main

import "fmt"

type book struct {
	price float64
	title string
}

func (b book) vat() float64 {
	return b.price * 9 / 100
}

func (b *book) discount() {
	b.price = b.price * 0.9
}

func main3() {
	b := book{
		price: 123.4,
		title: "Years of solitud",
	}

	vat := b.vat()
	fmt.Printf("vat is %v\n", vat)

	b.discount()
	fmt.Printf("Price after discount is %v\n", b.price)

	fmt.Printf("Book %#v", b)
}
