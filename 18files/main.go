package main

import (
	"fmt"
	"io"
	"os"
)

func main()  {
	fmt.Println("Welcome to files in golang!!")

	fileContent := "This is a file created using Golang!!"

	file, err := os.Create("./file.txt")
	checkNilErr(err)

	fileLength, err := io.WriteString(file, fileContent)
	checkNilErr(err)

	fmt.Println("File length:", fileLength)

	defer file.Close()

	readFile("./file.txt")
}

func readFile(filename string)  {
	fileBytes, err := os.ReadFile(filename)
	checkNilErr(err)

	fileContentStr := string(fileBytes)

	fmt.Println("File content:", fileContentStr)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
