package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("h", "There", "give the name you want to say hello")
	flag.Parse()
	if name != nil {
		finalString := "Hello " + *name
		fmt.Println(finalString)
		return
	}
	fmt.Println("No name provided")
}
