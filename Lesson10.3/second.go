package second

import "fmt"

func init() {
	fmt.Println("Package second is connected")
}

func Hello() string {
	return "Hello, Louis!"
}

/*

GOROOT=C:\Program Files\Go
ls "C:\Program Files\Go\src"
mkdir "C:\Program Files\Go\src\second"
cp ./Lesson10.3/second.go "C:\Program Files\Go\src\second"
go install second

*/
