package main

import "fmt"

type cari struct {
	adi           string
	soyadi        string
	alacak        float32
	borc          float32
	bakiye        float32
	bakiyeHesapla func(float32, float32) float32
}

func main() {
	c1 := cari{
		adi:           "Ali",
		soyadi:        "Meler",
		alacak:        1000,
		borc:          760,
		bakiyeHesapla: hesapla,
	}
	c1.bakiye = c1.bakiyeHesapla(c1.alacak, c1.borc)
	fmt.Println(c1.bakiye)

	c2Iskonto := cari{
		adi:           "Ahmet",
		soyadi:        "Medet",
		alacak:        1000,
		borc:          760,
		bakiyeHesapla: hesaplaIskontolu,
	}
	c2Iskonto.bakiye = c2Iskonto.bakiyeHesapla(c2Iskonto.alacak, c2Iskonto.borc)
	fmt.Println(c2Iskonto.bakiye)
}

func hesapla(a float32, b float32) float32 {
	return a - b
}

func hesaplaIskontolu(a float32, b float32) float32 {
	return a - b - 150
}
