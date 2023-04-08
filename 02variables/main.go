package main

import "fmt"

// It works outside main function
// var myVariable int = 10

// It doesn't work outside main function
// myVariable := 10

// FYI: If the first letter is a capital letter, Go will interpret the constant as public
const LoginToken = "login-token-123"

func main()  {
	var stringValue string = "hi"
	fmt.Println(stringValue)
	fmt.Printf("Variable is of type: %T \n", stringValue)
	fmt.Println("---")

	var boolValue bool = false
	fmt.Println(boolValue)
	fmt.Printf("Variable is of type: %T \n", boolValue)
	fmt.Println("---")

	var smallIntValue uint8 = 255
	fmt.Println(smallIntValue)
	fmt.Printf("Variable is of type: %T \n", smallIntValue)
	fmt.Println("---")

	var smallFloatValue float32 = 255.5123871239812389
	fmt.Println(smallFloatValue)
	fmt.Printf("Variable is of type: %T \n", smallFloatValue)
	fmt.Println("---")

	var bigFloatValue float64 = 255.5123871239812389
	fmt.Println(bigFloatValue)
	fmt.Printf("Variable is of type: %T \n", bigFloatValue)
	fmt.Println("---")

	// Default values
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)
	fmt.Println("---")

	// Implict type
	var myWebsite = "https://emmanuelbergmann.me"
	fmt.Println(myWebsite)
	fmt.Printf("Variable is of type: %T \n", myWebsite)
	fmt.Println("---")

	// No "var" style
	totalUsers := 300000
	fmt.Println(totalUsers)
	fmt.Printf("Variable is of type: %T \n", totalUsers)
	fmt.Println("---")

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
	fmt.Println("---")
}
