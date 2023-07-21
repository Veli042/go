package main

import "fmt"

type ogrenci struct {
	adi    string
	soyadi string
	numara int
}

func main() {

	var ogrenci1 ogrenci
	var ogrenci2 ogrenci = ogrenci{"Veli", "Yıkılmaz", 123}
	ogrenci3 := ogrenci{"Hamdi", "Açık", 234}

	ogrenci4 := ogrenci{numara: 42, adi: "Ahmet"}

	fmt.Println(ogrenci1)
	fmt.Println(ogrenci2.adi)
	fmt.Println(ogrenci3.soyadi)
	fmt.Println(ogrenci4.numara)
}
