package main

import "fmt"

func main() {

	fmt.Println("Задача 7.1")
	arrStr1 := [5]string{"Во", "поле", "береза", "стояла", "!"}
	fmt.Println(arrStr1)
	fmt.Println()

	fmt.Println("Задача 7.2")
	arrStr2 := [...]string{"яблоко", "груша", "слива", "абрикос"}
	fmt.Println(arrStr2)
	fmt.Println()

	fmt.Println("Задача 7.3")
	arrStr3 := [...]string{"яблоко", "груша", "помидор", "абрикос"}
	fmt.Println(arrStr3)
	arrStr3[2] = "персик"
	fmt.Println(arrStr3)
	fmt.Println()

	fmt.Println("Задача 7.4")
	arrInt1 := []int{5, 2, 8, 3, 1, 9}
	fmt.Println(arrInt1)
	fmt.Println()

	fmt.Println("Задача 7.5")
	arrEmpty := make([]int, 0, 10)
	fmt.Println("пустой срез:", arrEmpty, "длина:", len(arrEmpty), "ёмкость:", cap(arrEmpty))
	fmt.Println()

	fmt.Println("Задача 7.6")
	arrInt2 := make([]int, 0, 10)
	fmt.Println("срез до:", arrInt2, "длина:", len(arrInt2), "ёмкость:", cap(arrInt2))
	arrInt2 = append(arrInt2, 4, 1, 8, 9)
	fmt.Println("срез после:", arrInt2, "длина:", len(arrInt2), "ёмкость:", cap(arrInt2))
	fmt.Println()

	fmt.Println("Задача 7.7")
	slInt1 := []int{1, 2, 3}
	slInt2 := []int{4, 5, 6}
	slInt3 := []int{}
	slInt3 = append(slInt1, slInt2...)
	fmt.Println("срез после объединения:", slInt3)
	fmt.Println()

	fmt.Println("Задача 7.8")
	slInt8 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("срез до:", slInt8, "длина:", len(slInt8), "ёмкость:", cap(slInt8))
	slInt8 = append(slInt8[:3], slInt8[4:]...)
	fmt.Println("срез после:", slInt8, "длина:", len(slInt8), "ёмкость:", cap(slInt8))
	fmt.Println()
	fmt.Println()
	fmt.Println()

}
