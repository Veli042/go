package main

import "fmt"

const (
	version = "1.0.0"
	PI      = 3.14 //const pı sayısı gibi , e sayısı gibi değeri kesin olan değişmeyen ifadeleri tanımlamada kullanılır.
	e       = 2.71
)

func main() {
	fmt.Println("proje", version)
	version := "1.0.1"
	fmt.Println("proje", version)

}
