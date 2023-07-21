package main

func main() {

	/*var yaşlar [4]int
	yaşlar[0] = 10
	yaşlar[1] = 8
	yaşlar[2] = 18
	yaşlar[3] = 16

	for i := 0; i < 4; i++ {
		fmt.Println(yaşlar[i])
	}*/

	/*var adetler [5]int
	adetler[0] = 13
	adetler[1] = 16
	adetler[2] = 14
	adetler[3] = 15
	adetler[4] = 17

	for i := 0; i < len(adetler); i++ {
		fmt.Println(adetler[i])

	}*/

	/*var adetler [5]int
	adetler[0] = 12
	adetler[1] = 13
	adetler[2] = 14
	adetler[3] = 15
	adetler[4] = 16

	for i := 4; i > -1; i-- {
		fmt.Println(adetler[i])
	}*/

	/*var ogrenciSayisi int = 4
	var ogrenciAdlari [4]string = [4]string{"Ali", "Veli", "Ahmet", "Mehmet"}
	var ogrenciSoyadlari [4]string = [4]string{"Tan", "Yücel", "Ok", "Vural"}
	var ogrenciNotlari [4]int = [4]int{76, 40, 70, 55}

	for i := 0; i < ogrenciSayisi; i++ {
		if ogrenciNotlari[i] > 60 {
			fmt.Println(ogrenciAdlari[i], "", ogrenciSoyadlari[i], "Dersi geçtiniz.")
		} else {
			fmt.Println(ogrenciAdlari[i], "", ogrenciSoyadlari[i], "Dersten kaldınız.")
		}
	}*/

	/*var rakamlar [10]int = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(rakamlar); i++ {
		if i%2 == 0 {
			fmt.Println(rakamlar[i], "çifttir.")
		} else {
			fmt.Println(rakamlar[i], "tektir.")
		}
	}*/

	/*for i := 0; i < 100; i++ {
		fmt.Println(i)
		if i%2 == 0 && i%3 == 0 && i%5 == 0 {
			fmt.Println("2 3 ve 5 e tam bölünür")
		} else {
			fmt.Println("2 3 ve 5 e tam bölünmez")
		}


	}*/

	/*for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		} else {
			continue
		}
	}*/

	/*var toplam int = 0

	for i := 0; i < 11; i++ {
		toplam += i
	}
	fmt.Println(toplam)*/

	/*var toplam int = 0
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			toplam += i
			continue
		} else {
			continue
		}
	}
	fmt.Println(toplam)*/

	/*var n int
	var num1 int = 0
	var num2 int = 0
	var nextNumber int
	var count int

	fmt.Println("Terim sayısını giriniz:")
	fmt.Scan(&n)

	for count := 1; count <= n; count++ {
		fmt.Println(nextNumber, "")
		count += 1
		num1 = num2
		num2 = nextNumber
		nextNumber = num1 + num2
	}*/

	//GİRDİĞİMİZ SAYININ RAKAMLARI TOPLAMINI VEREN KOD

	/*var toplam int
	var n int
	fmt.Println("Bir Sayı giriniz:")
	fmt.Scan(&n)

	for i := 0; i < n+1; i++ {
		toplam += i
		continue
	}
	fmt.Println(toplam)*/

	//GİRDİĞİMİZ SAYIDA SADECE TEK SAYILARI YAZDIRAN KOD

	/*var n int
	fmt.Println("Bir sayı giriniz:")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			continue
		} else {
			fmt.Println(i)
		}
	}*/

	// GİRDİĞİMİZ SAYININ MÜKEMMEL SAYI OLUP OLMADIĞINI GÖSTEREN KOD

	/*var n int
	var toplam = 0
	fmt.Println("Bir sayı giriniz:")
	fmt.Scan(&n)

	for i := 1; i < n; i++ {
		if n%i == 0 {
			toplam += i
		}
	}
	if n == toplam {
		fmt.Println("Mükemmel Sayıdır.")
	} else {
		fmt.Println("Mükemmel sayı değildir.")
	}*/

	// GİRDİĞİMİZ SAYININ ASAL SAYI OLUP OLMADIĞINI GÖSTEREN KOD

	/*var n int

	fmt.Println("Bir sayı giriniz:")
	fmt.Scan(&n)

	if n > 2 {
		for i := 2; i < n; i++ {
			if n%i == 0 {
				fmt.Println(n, "Asal sayı değildir.")
				break
			} else {
				fmt.Println(n, "Asal sayıdır.")
				break
			}

		}

	} else if n == 2 {
		fmt.Println(n, "Asal sayıdır.")
	} else {
		fmt.Println(n, "Asal sayı değildir.")
	}*/

}
