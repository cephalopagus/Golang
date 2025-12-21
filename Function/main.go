package main

import "fmt"

func main() {

	name := "Dastan Bakaev"
	smth(name)
	num := float32(rtrn(1512))
	fmt.Print(num)
}

func smth(name string) {
	fmt.Println("Вывод из метода! Привет,", name)
}

func rtrn(i int) int {
	s := 12 * i
	fmt.Print("frfsfdfsdfsf")
	return s

}
