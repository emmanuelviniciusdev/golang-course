package main

import "fmt"

func main()  {
	fmt.Println("Methods class!!")

	user1 := User{
		Name: "Steven Paul Jobs", 
		Email: "steve.jobs@icloud.com", 
		Status: true, 
		DateOfBirth: "1955-02-24",
	}

	fmt.Println(user1)
	fmt.Println(user1.Name)
	fmt.Println(user1.Email)
	fmt.Println(user1.GetGreetings())
}

type User struct {
	Name string
	Email string
	Status bool
	DateOfBirth string
}

func (user User) GetGreetings() string {
	return "Hi, I'm " + user.Name + " and I was born in " + user.DateOfBirth + ". Nice 2 meet u :B"
}
