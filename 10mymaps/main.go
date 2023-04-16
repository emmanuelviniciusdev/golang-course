package main

import "fmt"

func main()  {
	fmt.Println("Maps in Golang!!")

	programmingLanguagesExtensions := make(map[string]string)

	programmingLanguagesExtensions["go"] = "Golang"
	programmingLanguagesExtensions["js"] = "Javascript"
	programmingLanguagesExtensions["py"] = "Python"
	programmingLanguagesExtensions["java"] = "Java"

	fmt.Println(programmingLanguagesExtensions)
	fmt.Println(programmingLanguagesExtensions["go"])
	fmt.Println(programmingLanguagesExtensions["java"])

	delete(programmingLanguagesExtensions, "java")

	fmt.Println(programmingLanguagesExtensions)

	for key, value := range programmingLanguagesExtensions {
		fmt.Println(key, value)
	}
}
