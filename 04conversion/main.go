package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	fmt.Println("Welcome to our salad app!!")
	fmt.Println("Please, rate our salad between 1 and 5:")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	numRating, err := strconv.ParseFloat(input, 64)

	if err != nil {
		fmt.Println("SOMETHING WENT WRONG")
		panic(err)
	}

	if numRating < 1 || numRating > 5 {
		panic("YOUR RATING IS INVALID!!!!!! PLEASE, RATE OUR SALAD CORRECTLY")
	}

	fmt.Println("Thanks for ur rating :):", numRating)

	numRating += 1

	fmt.Println("Oh, we've added 1 to your rating...")
	fmt.Println("Now your rating is:", numRating)
}
