package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Welcome to our time study lab!!")

	presentTime := time.Now()
	fmt.Println("presentTime", presentTime)

	formatedPresentTime := presentTime.Format("01-02-2006 15:04:05 Monday")
	fmt.Println("formatedPresentTime", formatedPresentTime)

	createdTime := time.Date(2010, time.October, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println("createdTime", createdTime)

	formatedCreatedTime := createdTime.Format("2006-01-02 Monday 15:04:05")
	fmt.Println("formatedCreatedTime", formatedCreatedTime)
}
