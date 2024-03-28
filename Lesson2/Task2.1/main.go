package main

import "fmt"

func main() {

	var mySym1 rune
	var myByte1 byte

	fmt.Printf("тип: %t, значение: %s", mySym1, mySym1)
	fmt.Println()
	fmt.Printf("тип: %t, значение: %s", myByte1, myByte1)

}

/*
  оба типа имеют значение по умолчанию = 0, т.к. оба являются числовыми типами.
*/
