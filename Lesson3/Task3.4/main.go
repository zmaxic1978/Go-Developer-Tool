package main

import "fmt"

const (
	myConst1 = 2 ^ iota
	myConst2
	myConst3
	myConst4
	myConst5
)

func main() {

	fmt.Println("Значения 5-ти глобальных констант:", myConst1, myConst2, myConst3, myConst4, myConst5)
}
