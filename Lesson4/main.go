package main

import "fmt"

func main() {

	fmt.Println("Task 4.1.")
	hello()

	fmt.Println("\nTask 4.2.")
	func() { fmt.Println("Hello, Go!") }()

	fmt.Println("\nTask 4.3.")
	var f = func() { fmt.Println("Hello, Go!") }
	hello3(f)

	fmt.Println("\nTask 4.4.")
	hello4()()

	fmt.Println("\nTask 4.5.")
	defer fmt.Println("Завершение работы")
	hello5()

	fmt.Println("\nTask 4.6. Первые 23 числа последовательности Фибоначчи")
	fmt.Printf("0 ")
	fmt.Printf("1 ")
	fibonachi(2, 22, 0, 1)
	fmt.Println()
	fmt.Println()
}

func hello() { fmt.Println("Hello, Go!") }

func hello3(fn func()) { fn() }

func hello4() func() {
	var f = func() { fmt.Println("Hello, Go!") }
	return f
}

func hello5() {
	fmt.Println("Hello, Go!")
}

func fibonachi(indMin int, indMax int, fprepre int, fpre int) {
	if indMin > indMax {
		return
	}
	var fcur = fprepre + fpre
	fmt.Printf("%d ", fcur)
	fibonachi(indMin+1, indMax, fpre, fcur)
}
