package main

import (
	cryptoRand "crypto/rand"
	"fmt"
	"math/big"
	mathRand "math/rand"
)

func main()  {
	var myNumberOne int = 10
	var myNumberTwo float64 = 10.5
	
	fmt.Println("The sum is:", myNumberOne + int(myNumberTwo))

	mathRandomNumber := mathRand.Intn(10)

	fmt.Println("Random number using math/random:", mathRandomNumber)

	cryptoRandomNumber, _ := cryptoRand.Int(cryptoRand.Reader, big.NewInt(10))

	fmt.Println("Random number using crypto/rand:", cryptoRandomNumber)
}
