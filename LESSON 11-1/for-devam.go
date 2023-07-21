package main

import "fmt"

func main() {
	var iller map[string]string
	iller = make(map[string]string)
	iller["01"] = "Adana"
	fmt.Println(iller)

	iller["01"] = "Yeni İl"
	fmt.Println(iller)

	il01, ok := iller["01"]
	fmt.Println(il01)

	il07 := iller["07"]
	if ok {
		fmt.Println(il07)
	} else {
		fmt.Println("07 Bulunamadı.")
	}
	delete(iller, "01")
	il01, ok1 := iller["01"]
	if ok1 {
		fmt.Println(il01)
	} else {
		fmt.Println("01 Bulunamadı.")
	}
}
