package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Print(i*j, "\t")

		}
		fmt.Println()
	}
	for i := 0; i < 10; i++ {
		fmt.Print(i, "fewf\n")
		time.Sleep(500 * time.Millisecond)
	}
}
