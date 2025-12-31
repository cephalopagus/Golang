package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// hash_map := map[string]time.Time{
	// 	"Dastan": time.Date(2002, 9, 5, 0, 0, 0, 0, time.UTC),
	// 	"Vika":   time.Date(2005, 8, 15, 0, 0, 0, 0, time.UTC),
	// }
	// fmt.Println(hash_map)

	// new_map := map[int]int{
	// 	1: 231322,
	// 	2: 12312323,
	// 	3: 342414,
	// }
	// new_map[4] = 214214

	// for k, v := range new_map {
	// 	fmt.Println("Ключ:", k, "\tЗначение:", v)
	// }
	// a, ok := new_map[1]
	// if ok == true {
	// 	fmt.Println(a, ok, "\n")
	// }
	// delete(new_map, 2)
	// for k, v := range new_map {
	// 	fmt.Println("Ключ:", k, "\tЗначение:", v)
	// }
	names := [10]string{
		"Dastan",
		"Atai",
		"Isma",
		"Aidar",
		"Nurbek",
		"Timur",
		"Azamat",
		"Eldar",
		"Bekzat",
		"Arsen",
	}
	mapa := map[int]string{}
	for _, v := range names {
		mapa[rand.Intn(200-100+1)+100] = v
	}
	for k, v := range mapa {
		fmt.Println("Ключ:", k, "\tЗначение:", v)
	}
}
