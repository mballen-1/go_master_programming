package main

import "fmt"

func cube(n float64) float64{
	return n * n * n
}

func main1() {
	lol := cube(10)
	fmt.Println(lol)
}
