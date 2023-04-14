package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("Welcome to a class on slices!!")

	var fruitList = []string{"Apples", "Strawberries"}

	fmt.Printf("Type of fruitList: %T \n", fruitList)

	fruitList = append(fruitList, "Grapes", "Mangos", "Pears")
	fmt.Println("fruitList", fruitList)

	// Removes the first element from slice
	fruitList = fruitList[1:]
	fmt.Println("fruitList", fruitList)

	// Slicing slices
	fruitList = fruitList[2:4]
	fmt.Println("fruitList", fruitList)

	highScores := make([]int, 3)

	highScores[0] = 888
	highScores[1] = 444
	highScores[2] = 999

	// This will crash
	// highScores[3] = 111

	highScores = append(highScores, 222, 111)

	fmt.Println("highScores", highScores)

	sort.Ints(highScores)
	fmt.Println("sorted highScores", highScores)

	highScoresAreSorted := sort.IntsAreSorted(highScores)
	fmt.Println("highScores is sorted?", highScoresAreSorted)
}
