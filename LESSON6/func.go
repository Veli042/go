package main

import "fmt"

func main() {
	var sayi1 float32
	var sayi2 float32
	var sonuc1 float32
	var sonuc2 float32
	var sonuc3 float32
	var sonuc4 float32

	fmt.Println("1. Sayıyı giriniz:")
	fmt.Scan(&sayi1)

	fmt.Println("2. Sayıyı giriniz:")
	fmt.Scan(&sayi2)

	sonuc1, sonuc2, sonuc3, sonuc4 = tümİslemler(sayi1, sayi2)
	fmt.Println("Toplam:", sonuc1)
	fmt.Println("Fark:", sonuc2)
	fmt.Println("Çarpım:", sonuc3)
	fmt.Println("Kalan:", sonuc4)
}

func tümİslemler(x float32, y float32) (add float32, extraction float32, impact float32, divide float32) {

	add = x + y
	extraction = x - y
	impact = x * y
	divide = x / y
	return add, extraction, impact, divide
}
