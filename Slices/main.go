package main

type User struct {
	Name    string
	Rating  float64
	Premium bool
}

func main() {
	user := User{
		Name:    "Dastan",
		Rating:  5,
		Premium: true,
	}
	user2 := User{
		Name:    "Atai",
		Rating:  2,
		Premium: false,
	}
	user3 := User{
		Name:    "Isma",
		Rating:  3,
		Premium: false,
	}
	users := [3]User{user, user2, user3}

	// arr := [5]int{1, 2, 3, 4, 5}
	// arr[2] += 2
	// str := "dastan bakaev"
	// fmt.Println(len(str))
	// for i := 0; i < len(str); i++ {

	// }
}
