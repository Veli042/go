package main

import "fmt"

func main() {
	deneme()
}

func deneme() {

	fmt.Println("Deneme başladı.")

	defer fmt.Println("Defer 1")
	defer fmt.Println("Defer 2")
	defer fmt.Println("Defer 3")

	fmt.Println("Deneme bitti.")
}

// defer () Last in First Out = LİFO son giren ilk çıkar .
