package main

import "fmt"

func main() {

	var adet int = 16

	switch {
	case adet < 24:
		fmt.Println("ADET <24")
		fallthrough
	case adet < 34:
		fmt.Println("ADET <34")
	default:
		fmt.Println("ADET BİLİNMİYOR.")

	}
}
