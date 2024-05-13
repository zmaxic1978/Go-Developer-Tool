package main

import "fmt"

func main() {
	ch := make(chan int)
	stop := make(chan struct{})
	go func() {
		<-ch
		stop <- struct{}{}
	}()
	close(ch) // закрываем канал, из закрытого канала читать можно и дочерняя горутина отработает
	<-stop
	fmt.Println("happy end")
}
