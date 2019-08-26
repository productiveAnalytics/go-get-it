package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	go func(delayInSecs time.Duration) {
		defer waitGrp.Done()

		fmt.Println("Ready for sleeping for ", delayInSecs, " seconds")
		time.Sleep(delayInSecs * time.Second)
		fmt.Println("Done sleeping for ", delayInSecs, " seconds...")
	}(10)

	go func() {
		defer waitGrp.Done()

		fmt.Println("Done working directly")
	}()

	// Let main wait for 2 goroutines
	waitGrp.Wait()
}
