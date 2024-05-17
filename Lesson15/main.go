package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Meteo interface {
	ReadTemp() string
	ChangeTemp(v string)
}

type Weather struct {
	mu    sync.RWMutex
	value string
}

func (w *Weather) ReadTemp() string {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.value
}
func (w *Weather) ChangeTemp(v string) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.value = v
}

func main() {

	fmt.Println("Задача 15.1")
	wg := sync.WaitGroup{}
	for i := 1; i <= 5; i++ {
		num := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("горутина:", num)
		}()
	}
	wg.Wait()
	fmt.Println()

	fmt.Println("Задача 15.2")
	var cnt int32
	wg2 := sync.WaitGroup{}
	for i := 1; i <= 500; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			atomic.AddInt32(&cnt, 1)
		}()
	}
	wg2.Wait()
	fmt.Println("Значение счетчика:", cnt)
	fmt.Println()

	fmt.Println("Задача 15.3")
	var cnt2 int
	var mu sync.Mutex
	wg3 := sync.WaitGroup{}
	for i := 1; i <= 500; i++ {
		wg3.Add(1)
		go func() {
			defer wg3.Done()
			defer mu.Unlock()
			mu.Lock()
			cnt2++
		}()
	}
	wg3.Wait()
	fmt.Println("Значение счетчика:", cnt2)
	fmt.Println()

	fmt.Println("Задача 15.4")
	wg4 := sync.WaitGroup{}
	var ro sync.Once
	for i := 1; i <= 10; i++ {
		num := i
		wg4.Add(1)
		go func() {
			defer wg4.Done()
			ro.Do(start)
			fmt.Println("горутина:", num)
		}()
	}
	wg4.Wait()
	fmt.Println()

	fmt.Println("Задача 15.5")
	w := Weather{mu: sync.RWMutex{}, value: ""}
	var m Meteo = &w

	for i := 0; i < 500; i++ {
		go func() {
			t := i
			val := fmt.Sprintf("Сейчас: %d C'", t)
			m.ChangeTemp(val)
		}()
		go func() {
			m.ReadTemp()
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println(m.ReadTemp())
	fmt.Println()
}

func start() { fmt.Println("Старт!") }
