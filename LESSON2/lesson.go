package main

import "fmt"

func main() {

	/*	var isActive bool
		isActive = true

		if isActive {
			fmt.Println("koşul doğru")
		} else {
			fmt.Println("koşul yanlış")
		}
	 	fmt.Println("Bitti")*/

	var veliyasi = 20
	var mustafayasi = 21

	/*if veliyasi == mustafayasi {
		fmt.Println("Yaşları Aynı")
	} else {
		fmt.Println("Yaşları Farklı")
	}
	fmt.Println("Bitti")*/

	/*if veliyasi < 10 || mustafayasi < 30 {
		fmt.Println("Koşul Doğru")
	} else {
		fmt.Println("Koşul Yanlış")
	}*/

	if veliyasi < 10 && mustafayasi < 30 {
		fmt.Println("Koşul Doğrudur.")
	} else {
		fmt.Println("Koşul Yanlışdır.")
	}

}
