package main

import "fmt"

type User struct {
	Name    string
	Rating  float64
	Premium bool
}

func main() {

	users := []User{
		User{
			Name:    "Dastan",
			Rating:  5,
			Premium: true,
		},
		User{
			Name:    "Atai",
			Rating:  2,
			Premium: false,
		},
		User{
			Name:    "Isma",
			Rating:  3,
			Premium: false,
		},
	}

	users = append(users, User{Name: "Adis", Rating: 4, Premium: true})
	for i := range users {
		if users[i].Premium {
			users[i].Rating += 2.0
		}
	}
	for i := range users {
		fmt.Println(users[i])
	}
	// str := "dastan"
	// for v := range str {
	// 	fmt.Println(string(str[v]))
	// }
}

// arr := [5]int{1, 2, 3, 4, 5}
// arr[2] += 2
// str := "dastan bakaev"
// fmt.Println(len(str))
// for i := 0; i < len(str); i++ {

// }
