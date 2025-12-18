package main

import "fmt"

func main() {

	name := "Dastan Bakaev"
	smth(name)
}

func smth(name string) {
	fmt.Println("Вывод из метода! Привет,", name)
}
