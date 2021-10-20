package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(n string, wg *sync.WaitGroup) {
	fmt.Printf("Hello, %s!\n", n)
	time.Sleep(1 * time.Second)
	wg.Done()
}

func main1() {
	var wg sync.WaitGroup
	wg.Add(1)
	go sayHello("Mr. Wick", &wg)
	wg.Wait()
}