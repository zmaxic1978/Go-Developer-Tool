package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Задача 14.1")
	go func() {
		fmt.Println("Привет из дочерней горутины!")
	}()
	time.Sleep(1 * time.Second)
	fmt.Println()

	fmt.Println("Задача 14.2")
	strChan := make(chan string, 1)
	strChan <- "Привет, строковый канал!"
	fmt.Println(<-strChan)
	fmt.Println()

	fmt.Println("Задача 14.3")
	strChan3 := make(chan string, 4)
	strChan3 <- "Привет"
	strChan3 <- "буферизованный канал!"
	fmt.Println(<-strChan3)
	fmt.Println(<-strChan3)
	fmt.Println()

	fmt.Println("Задача 14.4")
	
}
