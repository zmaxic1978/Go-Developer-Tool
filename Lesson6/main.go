package main

import "fmt"

type contract struct {
	ID     int
	Number string
	Date   string
}

func main() {

	fmt.Println("Задача 6.1")
	var valCt contract = contract{ID: 1, Number: "#000A\\n101", Date: "2024-01-31"}
	fmt.Printf("{ID:%d Number:%s Date:%s}\n\n", valCt.ID, valCt.Number, valCt.Date)

	fmt.Println("Задача 6.2")
	type contract struct {
		ID     int
		Number string
		Date   string
	}
	var valCt2 contract = contract{ID: 1, Number: "#000A\\t101", Date: "2024-01-31"}
	fmt.Printf("{ID:%d Number:%s Date:%s}\n\n", valCt2.ID, valCt2.Number, valCt2.Date)

	fmt.Println("Задача 6.3")
	fmt.Println(valCt.Descr())
	fmt.Println()

	fmt.Println("Задача 6.4")
	type contacts struct {
		Addss string
		Phone string
	}
	type user struct {
		ID   int
		Name string
		contacts
	}
	type employee struct {
		ID   int
		Name string
		contacts
	}
	var u = user{ID: 1, Name: "Борискин А.Н.", contacts: contacts{Addss: "ул.Лопуховая, д4 кв6", Phone: "+79565888454"}}
	var e = employee{ID: 1, Name: "Коммиссаров Е.В", contacts: contacts{Addss: "Невский пр д12 кв2", Phone: "+79219895184"}}
	fmt.Println(u.Addss, u.Phone, e.Addss, e.Phone)
	fmt.Println()
	fmt.Println()
}

func (c contract) Descr() string {
	return fmt.Sprintf("Договор № %s от %s", c.Number, c.Date)
}
