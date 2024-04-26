package main

import (
	"errors"
	"fmt"
)

// кастомная ошибка
type myFirstError struct {
	message string
}

func (e myFirstError) Error() string {
	return e.message
}

func main() {

	fmt.Println("Задача 11.1")
	err := errors.New("ошибка1")
	err = fmt.Errorf("ошибка2:%w", err)
	err = fmt.Errorf("ошибка3:%w", err)
	fmt.Println(err, "\n")

	fmt.Println("Задача 11.2")
	erru := errors.Unwrap(err)
	fmt.Println(erru, "\n")

	fmt.Println("Задача 11.3")
	err1 := errors.New("ошибка1")
	err2 := fmt.Errorf("ошибка2:%w", err1)
	err3 := fmt.Errorf("ошибка3:%w", err2)
	fmt.Println(err3)
	fmt.Printf("Была ли \"%v\" в \"%v\": %t\n\n", err1, err3, errors.Is(err3, err1))

	fmt.Println("Задача 11.4")
	fmt.Println(err3)
	fmt.Printf("Была ли ошибка типа \"myFirstError\" в \"%v\": %t\n\n", err3, errors.As(err3, new(myFirstError)))

}
