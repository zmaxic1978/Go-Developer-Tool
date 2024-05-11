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

	fmt.Println("Задача 13.1")
	txt := `{"number":1,"landlord":"Остап Бендер","tenat":"Шура Балаганов"}`
	fmt.Println("Исходный текст: ", txt)
	c := contract{}
	err := json.Unmarshal([]byte(txt), &c)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("Распарсенный объект: %+v \n\n", c)
}
