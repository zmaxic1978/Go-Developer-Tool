package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type ctxKey string

func main() {

	fmt.Println("Задача 16.1")
	ctx1 := context.Background()
	ctx1, cancel1 := context.WithCancel(ctx1)
	defer cancel1()
	wg := sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		num := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx1.Done():
					fmt.Println("стоп горутина: ", num)
					return
				default:
					time.Sleep(time.Second)
				}
			}
		}()
	}
	cancel1()
	wg.Wait()
	fmt.Println()

	fmt.Println("Задача 16.2")
	ctx2 := context.Background()
	ctx2, cancel2 := context.WithTimeout(ctx2, 2*time.Second)
	defer cancel2()
	wg2 := sync.WaitGroup{}
	for i := 1; i <= 5; i++ {
		num := i
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			for {
				select {
				case <-ctx2.Done():
					fmt.Println("стоп горутина: ", num)
					return
				default:
					time.Sleep(time.Second)
				}
			}
		}()
	}
	wg2.Wait()
	fmt.Println()

	fmt.Println("Задача 16.3")
	ctx3 := context.Background()
	dt := time.Now().Add(2 * time.Second)
	ctx3, cancel3 := context.WithDeadline(ctx3, dt)
	defer cancel3()
	wg3 := sync.WaitGroup{}
	for i := 1; i <= 5; i++ {
		num := i
		wg3.Add(1)
		go func() {
			defer wg3.Done()
			for {
				select {
				case <-ctx3.Done():
					fmt.Println("стоп горутина: ", num)
					return
				default:
					time.Sleep(time.Second)
				}
			}
		}()
	}
	wg3.Wait()
	fmt.Println()

	fmt.Println("Задача 16.4")
	ctx4 := context.Background()
	var key1 ctxKey = "some key1"
	var key2 ctxKey = "some key2"
	ctx4 = context.WithValue(ctx4, key1, "some value1")
	ctx4 = context.WithValue(ctx4, key2, "some value2")
	showContext(ctx4)
	fmt.Println()

	fmt.Println("Задача 16.5")
	wg5 := sync.WaitGroup{}
	var flag int32 = 0
	for i := 1; i <= 10; i++ {
		num := i
		wg5.Add(1)
		go func() {
			defer wg5.Done()
			for {
				fmt.Println("сложные вычисления горутины:", num)
				flg := atomic.LoadInt32(&flag)
				if flg == 1 {
					break
				}
				time.Sleep(time.Second)
			}
			fmt.Println("stop горутина:", num)
		}()
	}
	// вредная горутина
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("ой, всё!")
		atomic.AddInt32(&flag, 1)
	}()
	wg5.Wait()
	fmt.Println()

}

func showContext(ctx context.Context) {
	var key1 ctxKey = "some key1"
	var key2 ctxKey = "some key2"
	fmt.Println(key1, ":", ctx.Value(key1))
	fmt.Println(key2, ":", ctx.Value(key2))
}
