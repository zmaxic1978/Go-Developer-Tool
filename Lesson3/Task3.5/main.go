package main

import "fmt"

func main() {

	const (
		myConst1 = 2 ^ iota
		myConst2
		myConst3
		myConst4
		myConst5
	)

	fmt.Println("Значения 5-ти локальных констант: ", myConst1, myConst2, myConst3, myConst4, myConst5)
}
