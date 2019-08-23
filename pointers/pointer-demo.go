package main

import "fmt"

func main() {
	// Declare pointer
	var orgName *string = new(string)
	*orgName = "ProductiveAnalytics"
	fmt.Println("Dereferencing Pointer: " + *orgName)

	fmt.Println()

	var newName string = "Prod8ctive"

	// address-of operator
	newOrgNamePtr := &newName
	fmt.Println("Pointer Address & Content :")
	fmt.Println(newOrgNamePtr, *newOrgNamePtr)

	newName = "Prod8ctive.Analytics"
	fmt.Println("Pointer Address (same) & Content (updated) :")
	fmt.Println(newOrgNamePtr, *newOrgNamePtr)
}
