package main

import "fmt"

func main() {

	fmt.Println("Задача 9.1")
	fmt.Printf(fruitMarket("апельсин"))
	fmt.Printf(fruitMarket("яблоки"))
	fmt.Printf(fruitMarket("сливы"))
	fmt.Printf(fruitMarket("груши"))
	fmt.Printf(fruitMarket("бегемот"))
	fmt.Println()

	fmt.Println("Задача 9.2")
	srez := []int{1, 2, 3}
	for i := 0; i < 3; i++ {
		fmt.Println("v1:", srez[i])
	BreakPoint:
		for j := 0; j < 3; j++ {
			fmt.Println("\tv2:", srez[j])
			for k := 0; k < 3; k++ {
				fmt.Println("\t\tv3:", srez[k])
				for m := 0; m < 3; m++ {
					fmt.Println("\t\t\tv4:", srez[m])
					if m == 1 {
						break BreakPoint
					}
				}
			}
		}
	}
	fmt.Println()

	fmt.Println("Задача 9.3")
	checkFood("груша")
	checkFood("яблоко")
	checkFood("апельсин")
	checkFood("тыква")
	checkFood("огурец")
	checkFood("помидор")
	checkFood("арбуз")
	fmt.Println()
}

func fruitMarket(fruitName string) string {
	map1 := map[string]int{
		"апельсин": 5,
		"яблоки":   3,
		"сливы":    1,
		"груши":    0}

	val, ok := map1[fruitName]
	if ok {
		return fmt.Sprintln(fruitName, ":", val)
	} else {
		return fmt.Sprintln(fruitName, ": отсутствует в списке")
	}
}

func checkFood(foodName string) {
	switch foodName {
	case "груша", "яблоко", "апельсин":
		fmt.Println(foodName, " - это фрукт")
	case "тыква", "огурец", "помидор":
		fmt.Println(foodName, " - это овощь")
	default:
		fmt.Println(foodName, " - что-то странное...")
	}
}
