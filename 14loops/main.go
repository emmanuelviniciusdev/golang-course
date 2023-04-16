package main

import "fmt"

func main()  {
	fmt.Println("Welcome to loops in golang!!")

	dayList := []string{"Monday", "Tuesday", "Wednesday"}

	for i := 0; i < len(dayList); i++ {
		fmt.Println(dayList[i])
	}

	fmt.Println("---")

	for i := range dayList {
		fmt.Println(dayList[i])
	}

	fmt.Println("---")

	for _, day := range dayList {
		fmt.Println(day)
	}

	fmt.Println("---")

	myIndex := 0

	for myIndex < len(dayList) {
		if myIndex == 2 {
			goto myLabel
		}

		fmt.Println(myIndex)
		myIndex++
	}

	myLabel:
		fmt.Println("This is strange but I like it")
}
