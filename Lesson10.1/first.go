package first

import "fmt"

func init() {
	fmt.Println("Package first is connected")
}

func Hello() string {
	return "Hello, Dolly!"
}

/*

GOROOT=C:\Program Files\Go
ls "C:\Program Files\Go\src"
mkdir "C:\Program Files\Go\src\first"
cp ./Lesson10.1/first.go "C:\Program Files\Go\src\first"
go install first

*/
