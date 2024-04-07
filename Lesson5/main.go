package main

import "fmt"

type square int

func main() {

	var anyStr string = "Какая-то строка"
	var pntStr *string = &anyStr

	fmt.Println("Задача 5.1 и 5.2")
	fmt.Printf("Значение: %s, адрес: %x\n\n", *pntStr, pntStr)

	fmt.Println("Задача 5.3")
	*pntStr = "Новая поменянная строка"
	fmt.Printf("Значение: %s, адрес: %x\n\n", *pntStr, pntStr)

	fmt.Println("Задача 5.4")
	var valNum1 int = 100
	var valNum2 int = 200
	var valNum3 int = 300
	var valNum4 int = 400
	fmt.Printf("Значение: %d, адрес: %x\n", valNum1, &valNum1)
	fmt.Printf("Значение: %d, адрес: %x\n", valNum2, &valNum2)
	fmt.Printf("Значение: %d, адрес: %x\n", valNum3, &valNum3)
	fmt.Printf("Значение: %d, адрес: %x\n\n", valNum4, &valNum4)
	// адреса увеличиваются, как правило на величину типа пееременной INT (в х64 системах это 8 байт)

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
