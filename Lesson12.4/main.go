package main

import (
	"errors"
	"fmt"
)

type Bird interface{ Sing() string }

type Duck struct{ voice string }

func (d *Duck) Sing() string { return d.voice }

func main() {
	var d *Duck
	song, err := Sing(d)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(song)
}

func Sing(b Bird) (string, error) {
	if b != nil {
		return b.Sing(), nil
	}
	return "", errors.New("Ошибка пения!")
}

/*
1. На строке 15 инициализировать переменную d: var d *Duck = &Duck{ voice: "Кря"}. В консоли выведется "Кря".
   Почему это помогает: потому что даже если d = nil, сравнение на строке 26 "b != nil" будет давать True (интерфейс устроен из 2-х частей: типа и
   указателя на объект если хотя бы одно из полей заполнено, сравнение на nil будет давать false), и при вызове b.Sing() на строке 27,
   происходит паника, т.к. реализации метода Sing() нет.
2. На строке 25 интерфейс Bird заменить на *Duck, и тогда в функцию Sing будет передаваться ссылка на объект. Паники не будет, но поскольку переменная d = nil,
   то функция Sing в консоли вернет ошибку "Ошибка пения!"
3. На строке 15 инициализировать переменную d: var d *Duck = &Duck{ voice: "Кря"} . На строке 25 интерфейс Bird заменить на *Duck, и тогда в функцию Sing будет
   передаваться ссылка на существующий объект. В консоли выведется "Кря"
4. Убрать на строках 12 и 15 указатель (*), работать прямо с типом. Объект d типа Duck будет инициализирован автоматически значениями по умолчанию. В этом случае
   переменная d неявно преобразуется в интерфейс Bird, и далее у интерфейса вызывается метод Sing(). В консоль выведется пусто, т.к. voice при автоматической
   инициализации = "". Паники нет. 

*/