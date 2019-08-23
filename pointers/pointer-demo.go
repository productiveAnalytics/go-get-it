package main

import "fmt"

func main() {
	// Declare pointer
	var orgName *string = new(string)
	*orgName = "ProductiveAnalytics"
	fmt.Println("Dereferencing Pointer: " + *orgName)
}
