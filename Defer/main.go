package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		fmt.Println("ferwsff")
		fmt.Println("ferwsff1212")
	}()

	smt()
	fmt.Println("1")

}

func smt() {
	defer func() {
		fmt.Println("Print from Function Something")
		for i := 0; i < 6; i++ {
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}
	}()
	defer func() {
		fmt.Println("Print from Function Something nomber 222222222")
	}()
	fmt.Println("чтчотчочтчочто")
}
