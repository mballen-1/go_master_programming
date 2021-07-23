package main

import (
	"log"
)

func HandleError(e error){
	if e != nil {
		log.Fatalf("Oops, it failed -> %v", e)
	}
}


/*func main() {
	file, err := os.Open(path)
	HandleError(err)
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	HandleError(err)
	fmt.Printf("The data is: %s\n", data)
}*/
