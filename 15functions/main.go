package main

import "fmt"

func main()  {
	fmt.Println("Welcome to functions in golang!!")

	greeter()

	func ()  {
		fmt.Println("Anonymous!! func!!")	
	}()

	result := sum(1, 1)
	fmt.Println("Result:", result)

	result, premiumCertificateMsg := premiumSum(1, 1, 1, 1, 1)
	fmt.Println("Result:", result, premiumCertificateMsg)
}

func premiumSum(valueList ...int) (int, string) {
	result := 0

	for _, value := range valueList {
		result += value
	}

	return result, "-> This is a result from a PREMIUM function 8)"
}

func sum(num1 int, num2 int) int {
	return num1 + num2
}

func greeter() {
	fmt.Println("Namstey from Golang")
}
