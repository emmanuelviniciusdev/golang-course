package main

import "fmt"

func main()  {
	fmt.Println("Welcome to arrays in golang!!")

	var fruitList [5]string

	fruitList[0] = "Apples"
	fruitList[3] = "Grapes"
	fruitList[4] = "Strawberries"

	fmt.Println("fruitList", fruitList)
	fmt.Println("len(fruitList)", len(fruitList))
}