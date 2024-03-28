package main

import "fmt"

const (
	myConst1 = 1 << iota
	myConst2
	myConst3
	myConst4
	myConst5
)

func main() {

	fmt.Println("Значения 5-ти констант степени 2^n:", myConst1, myConst2, myConst3, myConst4, myConst5)
}
