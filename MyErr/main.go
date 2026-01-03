package main

import "fmt"

func main() {
	check := 0
	num, err := 2 / check
	if err != nil {
		fmt.Errorf("math err - %w", err)
	} else {
		fmt.Println(num)
	}
}
