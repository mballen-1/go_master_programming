package main

import (
	"io/ioutil"
	"os"
)


func main() {
	_, err := os.Create(path)
	HandleError(err)

	bs := []byte("The Go gopher is an iconic mascot!")
	err = ioutil.WriteFile(path, bs, 0644)
	HandleError(err)
}
