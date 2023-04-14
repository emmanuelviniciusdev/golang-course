package main

import "fmt"

func main()  {
	fmt.Println("Welcome to a class on pointers!!")

	myNumber := 23
	myNumberPtr := &myNumber

	fmt.Println("myNumber", myNumber)
	fmt.Println("myNumberPtr", myNumberPtr)
	fmt.Println("myNumberPtr (value)", *myNumberPtr)

	fmt.Println("Multiplying pointers' value...")

	*myNumberPtr = *myNumberPtr * 2

	fmt.Println("myNumber", myNumber)
	fmt.Println("myNumberPtr", myNumberPtr)
	fmt.Println("myNumberPtr (value)", *myNumberPtr)
}
