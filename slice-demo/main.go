package main

import "fmt"

func main() {

	my_arr := [5]int{1, 2, 3, 4, 5}
	fmt.Print("my arr (orig) : ")
	fmt.Println(my_arr)

	my_slice_from_arr := my_arr[:]

	fmt.Println("Updating last elem of Slice to be 55...Underlying array will be updated too.")
	my_slice_from_arr[4] = 55
	fmt.Print("my array (updated) : ")
	fmt.Println(my_arr)
	fmt.Print("my slice from array : ")
	fmt.Println(my_slice_from_arr)

	my_slice_extensible := []int{11, 222, 333, 4444, 55555}
	fmt.Print("my slice extensible (orig) : ")
	fmt.Println(my_slice_extensible)
	// my_slice_extensible[7] = 7 	// This will result in error
	my_slice_extensible = append(my_slice_extensible, 666)
	fmt.Print("my slice extensible (updated) : ")
	fmt.Println(my_slice_extensible)
}
