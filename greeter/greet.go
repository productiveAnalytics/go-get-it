package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "-", "Name to greet.")
	flag.Parse() // read parameters

	if *name == "-" {
		// generic greeting
		fmt.Println("Hi ProductiveAnalytics!")
	} else {
		// specific greeting
		fmt.Println("Hello, " + *name)
	}
}
