package main

import "fmt"

func main() {
	var vize float32
	var final float32
	var ortalama float32

	fmt.Println("Vize notunuzu giriniz:")
	fmt.Scan(&vize)

	fmt.Println("Final notunuzu giriniz:")
	fmt.Scan(&final)

	ortalama = vize*0.4 + final*0.6
	fmt.Println("ortalamaniz", ortalama)
	if ortalama < 60 {
		fmt.Println("Kaldiniz.")
	} else {
		fmt.Println("Geçtiniz")
	}
}
