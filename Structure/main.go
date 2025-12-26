package main

import "fmt"

type User struct {
	Name        string
	Age         int
	PhoneNumber string
	IsClosed    bool
	Rating      float64
}
type Company struct {
	Name    string
	Ceo     string
	Address string
}

func (c *Company) Fired(num int) {
	fmt.Println("ТЫ УВОЛЕН!!!!!!!! ИЗ", c.Name)
	c.Name = "Mdigital"
	fmt.Println(num)
}

func main() {
	/*user := User{
		Name:        "Dastan",
		Age:         23,
		PhoneNumber: "0773428201",
		IsClosed:    true,
		Rating:      5.5,
	}
	fmt.Println(user.Name)
	user.Name = "Isma"
	user.hello()*/

	company := Company{
		Name:    "Technopark",
		Address: "Gorkogo 1/2",
		Ceo:     "Anvar",
	}
	num := 12
	company.Fired(num)
	fmt.Println(company.Name)

}

func (u User) hello() {
	fmt.Println("my name is", u.Name)

}
