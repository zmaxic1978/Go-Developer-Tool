package main

import (
	"fmt"

	v100 "github.com/zmaxic1978/trans_mod/100"
	v110 "github.com/zmaxic1978/trans_mod/110"
	v200 "github.com/zmaxic1978/trans_mod/v2/200"
	v210 "github.com/zmaxic1978/trans_mod/v2/210"
)

const Path = `R:\Docs\Работа\Go-Lessons\Go-Developer-Tool\Lesson13.5\`

func main() {

	err := v100.Do(Path+"patients", Path+"result10")
	if err != nil {
		fmt.Println("Ошибка выполнения trans_mod v1.0.0: ", err)
	}

	err = v110.Do(Path+"patients", Path+"result11")
	if err != nil {
		fmt.Println("Ошибка выполнения trans_mod v1.1.0: ", err)
	}

	err = v200.Do(Path+"patients", Path+"result20")
	if err != nil {
		fmt.Println("Ошибка выполнения trans_mod v2.0.0: ", err)
	}

	err = v210.Do(Path+"patients", Path+"result21")
	if err != nil {
		fmt.Println("Ошибка выполнения trans_mod v2.1.0: ", err)
	}
}
