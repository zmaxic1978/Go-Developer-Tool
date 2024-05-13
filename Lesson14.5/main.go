package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	ch := make(chan int, 1)
	stop := make(chan struct{}, 2)

	go func() {
	OUT:
		for {
			select {
			case <-stop:
				break OUT
			case v, ok := <-ch:
				if !ok {
					break OUT
				}
				fmt.Println(v)
			default:
				continue
			}
		}
		fmt.Println("завершение работы горутины_1")
	}()

	go func() {
		var i int
	OUT:
		for {
			i++
			select {
			case <-stop:
				break OUT
			default:
				time.Sleep(time.Second)
				if ch == nil {
					continue
				}
				ch <- i
			}
		}
		fmt.Println("завершение работы горутины_2")
	}()

	//Способ 1. можно вычитать данные из канала раньше другой роутины (проверено, работает. но не гарантирует на 100%)
	//go func() { for { <-ch } }()

	//Способ 2. блокируем канал stdout от вывода на время работы горутин, потом восстанавливаем.
	oldStdout := os.Stdout
	os.Stdout = nil

	time.Sleep(5 * time.Second)
	os.Stdout = oldStdout

	stop <- struct{}{}
	stop <- struct{}{}
	time.Sleep(time.Second)
	fmt.Println("завершение работы главной горутины")
}
