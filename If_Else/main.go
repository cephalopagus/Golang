package main

import (
	"fmt"
)

func main() {
	num := 11

	/*if num >= 10 || num <= 5 {
		fmt.Print("говно")
	} else {
		fmt.Print("норм")
	}*/

	switch num {
	case 8:
		fmt.Println("Заебись")
	case 10:
		fmt.Println("Говно")
	default:
		fmt.Println("НИХУЯ НЕТ!")
	}

}
