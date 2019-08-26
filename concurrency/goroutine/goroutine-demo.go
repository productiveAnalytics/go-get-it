package main

import (
	"fmt"
	"time"
)

func myFunc(from string, iterations int) {
	for i := 0; i < iterations; i++ {
		time.Sleep(5 * time.Second)
		fmt.Println(from, ": ", i)
	}
}

func main() {
	fmt.Println("Anonymous function GoRoutine execution...")
	go func(from string, iterations int) {
		for i := 0; i < iterations; i++ {
			time.Sleep(5 * time.Second)
			fmt.Println(from, ": ", i)
		}
	}("anonymous-goroutine", 10)

	fmt.Println("Direct execution...")
	myFunc("direct", 5)

	fmt.Println("GoRoutine execution...")
	go myFunc("goroutine", 5)

	time.Sleep(1 * time.Second)

	fmt.Println("Press Enter to finish...")
	fmt.Scanln()
	fmt.Print("Done!")
}
