package main

import "fmt"

type user struct {
	Id        int
	FirstName string
	LastName  string
}

func main() {
	// User 1 is being initialized by populating attributes of struct
	var user1 user
	user1.Id = 1
	user1.FirstName = "Arther"
	user1.LastName = "Zooie"
	fmt.Print("User1 : ")
	fmt.Println(user1)

	// User 2 being intialized implicit way
	user2 := user{Id: 2,
		FirstName: "Zelda",
		LastName:  "Quincy", // Note: for multi-line, must have , - otherwise compile would add ;
	}
	fmt.Print("User2 : ")
	fmt.Println(user2)
}
