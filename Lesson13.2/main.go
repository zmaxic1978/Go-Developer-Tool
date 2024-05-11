package main

import (
	"encoding/json"
	"fmt"
)

type contract struct {
	Number   int
	Landlord string
	Tenat    string
}

func main() {

	fmt.Println("Задача 13.2")
	c := contract{Number: 2, Landlord: "Остап", Tenat: "Шура"}
	fmt.Printf("Исходный объект: %+v \n", c)
	res, err := json.Marshal(c)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("Собранный json: ", string(res))
	fmt.Println()
}
