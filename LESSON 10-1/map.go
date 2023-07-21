package main

import "fmt"

func main() {
	var iller map[string]string
	iller = make(map[string]string)
	if iller == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not")
	}

	iller["01"] = "Adana"
	iller["06"] = "Ankara"
	iller["07"] = "Antalya"
	iller["34"] = "İstanbul"
	iller["42"] = "Konya"
	fmt.Println(iller)

	var ögrenciler map[int]string
	ögrenciler = make(map[int]string)

	ögrenciler[56] = "Ahmet"
	ögrenciler[88] = "Mehmet"
	ögrenciler[70] = "Can"
	ögrenciler[65] = "Sefa"
	ögrenciler[40] = "Ali"
	fmt.Println(ögrenciler)
	fmt.Println(ögrenciler[88])
}
