package main

import "fmt"

type square int

func main() {

	var anyStr string = "Какая-то строка"
	var pntStr *string = &anyStr

	fmt.Println("Задача 5.1 и 5.2")
	fmt.Printf("Значение: %s, адрес: %x\n\n", *pntStr, pntStr)

	fmt.Println("Задача 5.3 и 5.4")
	*pntStr = "Новая поменянная строка"
	fmt.Printf("Значение: %s, адрес: %x\n\n", *pntStr, pntStr)

	fmt.Println("Задача 5.5")
	change(pntStr)
	fmt.Printf("Значение: %s, адрес: %x\n\n", *pntStr, pntStr)

	fmt.Println("Задача 5.6 и 5.7")
	var s square = 25
	fmt.Printf("Значение s: %d\n", s)
	s = 30
	s += 15
	fmt.Printf("Значение s (30 + 15): %d\n\n", s)

	fmt.Println("Задача 5.8")
	s = 34
	s += 10
	fmt.Printf("Значение s: %s\n", s.getStr())

}

func change(pntStr *string) {
	if pntStr != nil {
		*pntStr = "И вновь значение поменялось"
	}
}

func (val square) getStr() string {
	return fmt.Sprintf("%d m^2", val)
}
