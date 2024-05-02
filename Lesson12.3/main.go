package main

import "fmt"

type Doc interface{ Format() }

type Xml struct{}

func (x Xml) Format() { fmt.Println("Данные в формате xml") }

type Csv struct{}

func (c Csv) Format() { fmt.Println("Данные в формате csv") }

func main() {
	x := Xml{}
	Report(x)
	c := Csv{}
	Report(c)
}

func Report(x Doc) {
	x.Format()
}
