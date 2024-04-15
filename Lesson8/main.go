package main

import "fmt"

func main() {

	fmt.Println("Задача 8.1")
	map1 := map[string]struct{}{
		"слон":    struct{}{},
		"бегемот": struct{}{},
		"носорог": struct{}{},
		"лев":     struct{}{}}
	fmt.Println(map1)
	fmt.Println()

	fmt.Println("Задача 8.2")
	map2 := map[string]int{
		"слон":    3,
		"бегемот": 0,
		"носорог": 5,
		"лев":     1}
	m, ok := map2["слон"]
	fmt.Printf("Животное: %s, количество: %d (есть в карте: %v)\n", "cлон", m, ok)
	m, ok = map2["бегемот"]
	fmt.Printf("Животное: %s, количество: %d (есть в карте: %v)\n", "бегемот", m, ok)
	m, ok = map2["осьминог"]
	fmt.Printf("Животное: %s, количество: %d (есть в карте: %v)\n", "осьминог", m, ok)
	fmt.Println()

	fmt.Println("Задача 8.3")
	map3 := map[string]struct{}{
		"слон":    struct{}{},
		"бегемот": struct{}{},
		"носорог": struct{}{},
		"лев":     struct{}{}}
	fmt.Println("Карта до: ", map3)
	delete(map3, "бегемот")
	fmt.Println("Карта после: ", map3)
	fmt.Println()

	fmt.Println("Задача 8.4")
	map4 := map[string]struct{}{
		"слон":    struct{}{},
		"бегемот": struct{}{},
		"носорог": struct{}{},
		"лев":     struct{}{}}
	fmt.Println("Карта до: ", map4)
	map4["выдра"] = struct{}{}
	fmt.Println("Карта после: ", map4)
	fmt.Println()

	fmt.Println("Задача 8.5")
	map5 := map[string]int{
		"слон":    3,
		"бегемот": 0,
		"носорог": 5,
		"лев":     1}
	fmt.Println("Карта до: ", map5)
	map5["бегемот"] = 2
	fmt.Println("Карта после: ", map5)
	fmt.Println()
}
