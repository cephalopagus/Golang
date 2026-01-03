package main

import "fmt"

type Animal interface {
	makeSound()
}
type Cat struct{}
type Dog struct{}

func (c *Cat) makeSound() {
	fmt.Println("MEOW")
}
func (c *Dog) makeSound() {
	fmt.Println("BARK")
}
func main() {

	var c, d Animal = &Cat{}, &Dog{}
	c.makeSound()
	d.makeSound()
	Say(c)

}
func Say(a Animal) {
	a.makeSound()
}
