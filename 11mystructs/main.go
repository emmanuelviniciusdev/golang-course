package main

import "fmt"

func main()  {
	fmt.Println("Structs class!!")

	user1 := User{
		Name: "Steven Paul Jobs", 
		Email: "steve.jobs@icloud.com", 
		Status: true, 
		DateOfBirth: "1955-02-24",
	}

	fmt.Println(user1)
	fmt.Println(user1.Name)
	fmt.Println(user1.Email)
}

type User struct {
	Name string
	Email string
	Status bool
	DateOfBirth string
}
