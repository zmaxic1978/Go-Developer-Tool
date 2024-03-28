package main

import "fmt"

func main() {

	var myVal1 = 16
	var myVal2 = 3
	var myRes = myVal1 / myVal2
	var myRem = myVal1 % myVal2

	fmt.Printf("Результат: %d, остаток от деления: %d, тип результата: %t", myRes, myRem, myRem)
}
