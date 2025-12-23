package main

import "fmt"

func main() {
	number := 10

	pointer := &number
	metod(pointer)
	fmt.Println(number, "finish")

	var nil_ptr *int
	fmt.Println(nil_ptr) //нулевой указататель, ведет в никуда

	if nil_ptr != nil {
		fmt.Println("Все заебись!")
	} else {
		fmt.Println("Хуево")
	}

}

func metod(num *int) {
	fmt.Println(num)
	fmt.Println(*num)
	*num = 120
	fmt.Println(*num)
}
