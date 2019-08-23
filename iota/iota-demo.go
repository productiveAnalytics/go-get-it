package main

import "fmt"

// constant block #1
const (
	first  = iota
	second = iota
	third  = iota
	fourth = 4 << iota
	PI     = 3.14
)

// constant block #2
const (
	fifth   = iota // iota will reset to zero
	sixth          // No need to define, it will same as above i.e. iota
	seventh = iota + 10
)

func main() {
	fmt.Println("From constant block #1")
	fmt.Println(first, second, third, fourth, PI)

	fmt.Println("From constant block #2")
	fmt.Println(fifth, sixth, seventh)
}
