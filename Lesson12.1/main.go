package main

import "fmt"

func main() {
	a := 1
	do(a)
}

func do(v any) {
	a := 10	
	val, ok := v.(int)
	if ok { a += val } else { panic("ошибка увеличения А на Val") }
	fmt.Println(a)
}
