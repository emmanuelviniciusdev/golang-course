package main

import (
	"encoding/json"
	"fmt"
)

type MyStruct struct {
	Foo string `json:"fooName"`
	Bar int `json:"barName"`
	Baz []string `json:"bazName,omitempty"`
	SecretValueHiddenFromJSON float32 `json:"-"`
}

func main()  {
	fmt.Println("Welcome to JSON class!!")

	// EncodeJSON()

	DecodeJSON()
}

func DecodeJSON()  {
	myJsonBytes := []byte(`
		{
			"fooName": "Foo",
			"barName": 1,
			"bazName": ["Baz_1", "Baz_2"]
		}
	`)

	isJsonInvalid := !json.Valid(myJsonBytes)

	if isJsonInvalid {
		panic("THE JSON IS INVALID!!!!")
	}

	var myStruct MyStruct
	var myMap map[string]interface{}

	json.Unmarshal(myJsonBytes, &myStruct)
	json.Unmarshal(myJsonBytes, &myMap)
	
	fmt.Println("myStruct:")
	fmt.Println(myStruct)
	fmt.Printf("%#v\n", myStruct)

	fmt.Println("myMap:")
	fmt.Println(myMap)
	fmt.Printf("%#v\n", myMap)
}

func EncodeJSON()  {
	myStruct := []MyStruct{
		{"Test 1", 1, []string{"Hello", "World"}, 99.5},
		{"Test 2", 2, []string{"Hola", "Mundo"}, 99.5},
		{"Test 3", 3, []string{"Olá", "Mundo"}, 99.5},
		{"Test 4", 4, nil, 99.5},
	}

	myJson, err := json.MarshalIndent(myStruct, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Println(myStruct)
	fmt.Println(string(myJson))
}
