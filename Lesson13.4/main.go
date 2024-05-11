package main

import (
	"encoding/xml"
	"fmt"
)

type contracts struct {
	List []contract `xml:"contract"`
}

type contract struct {
	Number   int    `xml:"number"`
	Landlord string `xml:"landlord"`
	Tenat    string `xml:"tenat"`
}

func main() {
	fmt.Println("Задача 13.4")
	cs := contracts{}
	c := contract{Number: 1, Landlord: "Остап Бендер", Tenat: "Шура Балаганов"}
	cs.List = append(cs.List, c)
	fmt.Printf("Исходный объект: %+v \n\n", cs)
	res, err := xml.Marshal(cs)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("Собранный XML: ", xml.Header, string(res))
	fmt.Println()
}
