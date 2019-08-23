package main

import "fmt"

// constant block #1
const (
	first  = "1st"
	second = iota
	third  = iota
	fourth = 4 << iota
)

// constant block #2
const (
	fifth   = iota // iota will reset to zero
	sixth   = iota
	seventh = iota + 10
)

func main() {
	fmt.Println("From constant block #1")
	fmt.Println(first, second, third, fourth)

	fmt.Println("From constant block #2")
	fmt.Println(fifth, sixth, seventh)
}
