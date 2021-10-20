package main

import "fmt"

type vehicle interface {
	License() string
	Name() string
}

type car struct {
	licenceNo string
	brand     string
}

func (c car) License() string {
	return c.licenceNo
}

func (c car) Name() string {
	return c.brand
}

func main1() {
	var v vehicle = car{
		brand:     "Mercedez",
		licenceNo: "023948",
	}

	fmt.Printf("Licence is %v\n", v.License())
	fmt.Printf("Brand is %v\n", v.Name())
}
