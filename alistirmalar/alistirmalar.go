package main

import (
	"fmt"
)

func main() {

	fmt.Println("Veli Yıkılmaz")
	//  Girilen Sayının Pozitif, Negatif, ya da 0 Olduğunu Bulan Python Örneği

	/*var n float32

	fmt.Println("Bir sayı giriniz:")
	fmt.Scan(&n)

	if n > 0 {
		fmt.Println(n, "Pozitiftir.")
	} else if n < 0 {
		fmt.Println(n, "Negatiftir.")
	} else {
		fmt.Println(n, "Sıfırdır.")
	}*/

	// Yaşı Girilen Kişinin Ehliyet Alıp Alamayacağını Gösteren Python Örneği

	/*var n int

	fmt.Println("Yaşınızı giriniz:")
	fmt.Scan(&n)

	if n > 18 {
		fmt.Println("Ehliyet alabilirsiniz.")
	} else {
		fmt.Println("Ehliyet alamazsınız.")
	}*/

	//1-100 Arası 3′ e ve 5′ e tam bölünen sayıları bulan Python Örneği

	/*for i := 0; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i)
		}
	}*/

	//1 den Kullanıcının Girdiği Sayıya Kadar Sayıları Listeleyen Python Örneği

	/*var n int

	fmt.Println("Bir sayı giriniz:")
	fmt.Scan(&n)

	for i := 1; i < n; i++ {
		fmt.Println(i)
	}*/

	//Kenarları Girilen Dikdörtgenin Alanı ve Çevresini Bulan Python Örneği

	/*var KisaKenar float32
	var UzunKenar float32
	var Alan float32
	var cevre float32

	fmt.Println("kısa kenarı giriniz:")
	fmt.Scan(&KisaKenar)

	fmt.Println("uzun kenarı giriniz:")
	fmt.Scan(&UzunKenar)

	Alan = KisaKenar * UzunKenar
	fmt.Println("dikdörtgenin alanı:", Alan)

	cevre = (KisaKenar * 2) + (UzunKenar * 2)
	fmt.Println("dikdörtgenin çevresi:", cevre)*/

	//Kullanıcın girdiği iki sayı arasındaki sayıların toplamını gösteren Python Örneği.

	/*var n float32
	var k float32
	var toplam float32

	fmt.Println("bir sayı giriniz:")
	fmt.Scan(&n)

	fmt.Println("bir sayı giriniz:")
	fmt.Scan(&k)

	for i := n; i < k; i++ {
		toplam += i
	}
	fmt.Println("Arasındaki sayıların toplamı", toplam)*/

	//kullanıcının girmiş olduğu 2 sayı arasında olan tek
	//ve çift sayıların toplamını ayrı ayrı bulan ve sonucu ekranda gösteren Python Örneği

	/*var n int
	var k int
	var cifttoplam int
	var tektoplam int

	fmt.Println("bir sayı giriniz:")
	fmt.Scan(&n)

	fmt.Println("bir sayı giriniz:")
	fmt.Scan(&k)

	for i := n; i < k; i++ {
		if i%2 == 0 {
			cifttoplam += i
		} else {
			tektoplam += i
		}
	}
	fmt.Println("çift sayıların toplamı :", cifttoplam)
	fmt.Println("tek sayıların toplamı :", tektoplam)*/

	//Maaşı ve zam oranı girilen işçinin zamlı maaşını hesaplayarak ekranda gösteren Python örneği:

	/*var maaş float32
	var zam float32
	var yenimaaş float32

	fmt.Println("Maaşınızı giriniz:")
	fmt.Scan(&maaş)

	fmt.Println("zam oranını giriniz:")
	fmt.Scan(&zam)

	yenimaaş = maaş + (maaş)*(zam*(0.01))

	fmt.Println("yenimaaş:", yenimaaş)*/

	// yarıçapı verilen çemberin çevresi ve alanını veren kod

	/*var r int = 5
	var alan int
	var cevre int

	alan = r * r
	fmt.Println("Çemberin alanı:", alan)
	cevre = 2 * r * 3
	fmt.Println("Çemberin çevresisi:", cevre)*/

	//FAKTÖRİYEL HESAPLAMA

	/*var sayi int
	var deger int = 1

	fmt.Println("Faktöriyelini bulmak istediğiniz sayıyı giriniz:")
	fmt.Scan(&sayi)

	for i := 0; i < sayi; i++ {
		deger = deger * (i + 1)
	}
	fmt.Println("Faktöriyel:", deger)*/

	//OYUN

	/*var tuttugumSayi int = 42
	var n int

	fmt.Println("Bir sayı tuttum . Bil bakalım kaç:")
	fmt.Scan(&n)

	if n != tuttugumSayi {
		fmt.Println("Tekrar deneyiniz:")

	} else {
		fmt.Println("Tebrikler..")

	}*/

	/*var x int = 42

	for {
		r1 := rand.Intn(100)
		fmt.Println(r1)

		if r1 == x {
			fmt.Println(r1)
			break
		}

	}*/

	/*var r1 int = 0
	var toplam int = 0

	for r1 < 100 {
		r1 = rand.Intn(100)
		fmt.Println(r1)
		if r1%2 == 0 {
			continue
		} else {
			toplam += r1
			break
		}
	}
	fmt.Println(toplam)*/

	/*var n int
	var sonuc1 int
	var sonuc2 int
	var sonuc3 int

	fmt.Println("Bir sayı giriniz:")
	fmt.Scan(&n)

	for i := 1; i < n; i++ {
		sonuc3 += i
		if i%2 == 0 {
			sonuc1 += i
		} else if i%2 != 0 {
			sonuc2 += i
		}

	}
	fmt.Println("Çift sayıların toplamı:", sonuc1)
	fmt.Println("Tek sayıların toplamı:", sonuc2)
	fmt.Println("Tüm sayıların toplamı:", sonuc3)*/

	// GİRDİĞİMİZ BİR SAYININ ASAL OLUP OLMADIĞINI DÖNDÜREN KOD

	/*var n int

	fmt.Println("Bir sayı giriniz:")
	fmt.Scan(&n)

	if n > 2 {
		for i := 2; i < n; i++ {
			if n%i == 0 {
				fmt.Println("Sayı asal değildir.")
				break
			} else {
				fmt.Println("Sayı asaldır.")
				break
			}
		}
	} else if n == 2 {
		fmt.Println("Sayı asaldır.")
	} else {
		fmt.Println("Sayı asal değildir.")
	}*/

	//1 saate kadar: 10 TL
	//1-5 saat arası: Saat başı 8 TL
	//5 saatten fazla: Saat başı 6 TL
	//Buna göre kullanıcının girdiği otoparkta kalınan saat süresine göre ödenecek miktarı bularak ekrana yazdırınız

	/*var saat float32
	var ucret float32

	fmt.Println("Kaldığınız süreyi giriniz:")
	fmt.Scan(&saat)

	if saat <= 1 {
		ucret := 10 * saat
		fmt.Println(ucret, "TL ödemeniz gereken tutar.")
	} else if saat > 1 && saat < 6 {
		ucret = 8 * saat
		fmt.Println(ucret, "TL ödemeniz gereken tutar.")
	} else if saat > 5 {
		ucret = 6 * saat
		fmt.Println(ucret, "TL ödemeniz gereken tutar.")
	}*/

	//kilo aralıklarına göre zayıf ya da kilolu olduğunu söyleyen kod

	/*var kilo float32

	fmt.Println("Kilonuzu giriniz:")
	fmt.Scan(&kilo)

	if kilo < 55 {
		fmt.Println("Zayıfsınız.")
	} else if kilo > 55 && kilo < 80 {
		fmt.Println("Normal kilodasınız.")
	} else {
		fmt.Println("Kilolusunuz.")
	}*/

	/*var asitliİcecekler int = 0
	var madensulari int = 0
	var sular int = 0
	var dolapRafKoliİcecekAdetleri [4][2][3]int = [4][2][3]int{{{3, 5, 2}, {2, 3, 4}}, {{5, 7, 8}, {1, 2, 3}}, {{3, 4, 5}, {6, 5, 4}}, {{9, 8, 7}, {5, 4, 7}}}

	for rafNo := 0; rafNo < 4; rafNo++ {
		for koliNo := 0; koliNo < 2; koliNo++ {
			var asitliİcecek int = dolapRafKoliİcecekAdetleri[rafNo][koliNo][0]
			var madensuyu int = dolapRafKoliİcecekAdetleri[rafNo][koliNo][1]
			var su int = dolapRafKoliİcecekAdetleri[rafNo][koliNo][2]
			asitliİcecekler += asitliİcecek
			madensulari += madensuyu
			sular += su
		}
	}
	fmt.Println(asitliİcecekler)
	fmt.Println(madensulari)
	fmt.Println(sular)*/

	// armstrong sayısı
	/*var sayi int
	var toplam = 0

	fmt.Println("Bir sayı giriniz:")
	fmt.Scan(&sayi)

	basamak := string(sayi)

	for x := 0; x < len(basamak); x++ {
		rakam := (string(x)) * (len(basamak))
		toplam += rakam
		if sayi == toplam {
			fmt.Println("Bir armstrong sayısıdır.")
		} else {
			fmt.Println("Armstrong sayısı değildir.")
		}
	}*/

	// Bir dizi verildiğinde, dizide hangi elemandan kaç tane olduğunu veren bir algoritma
	//şeması hazırlayınız.

	//var rakamlar [10]string = [10]string{"3", "4", "5", "3", "4", "5", "5", "4", "5", "4"}

	//var kucukHarfler [29]string = [29]string{"a", "b", "c", "ç", "d", "e", "f", "g", "ğ", "h", "ı", "i", "j", "k", "l", "m", "n", "o", "ö", "p", "r", "s", "ş", "t", "u", "ü", "v", "y", "z"}

	//var buyukHarfler [29]string = [29]string{"A", "B", "C", "Ç", "D", "E", "F", "G", "Ğ", "H", "I", "İ", "J", "K", "L", "M", "N", "O", "Ö", "P", "R", "S", "Ş", "T", "U", "Ü", "V", "Y", "Z"}

	/*var sayilar [10]int = [10]int{3, 4, 5, 5, 4, 3, 3, 3, 4, 4}

	fmt.Println("Dizideki en küçük sayı:", sayilar[0])
	fmt.Println("Dizideki en büyük sayı:", sayilar[2])

	for i := 0; i < 10; i++ {
		var bakilacakSayi int = int(sayilar[i])
		if i == bakilacakSayi {
			fmt.Println("3 rakamı dizinin içinde", bakilacakSayi, "adet var.")
		}
	}*/
	/*var sayilar [10]float32 = [10]float32{9, 3, 2, 5, 9, 9, 4, 5, 7, 9}
	var toplam float32
	var toplam2 float32
	for i := 0; i < len(sayilar); i++ {

		toplam += sayilar[i]
		toplam2 = toplam / float32(len(sayilar))

	}
	fmt.Println("dizideki sayıların toplamı:", toplam)
	fmt.Println("dizideki sayıların ortalaması:", toplam2)*/

	/*var dizi []int = []int{}
		var yeniNum int
		for {
			fmt.Println("Sayı giriniz(Çıkış için 0'ı tuşlayınız):")
			fmt.Scan(&yeniNum)
			if yeniNum == 0 {
				break
			}
			dizi = append(dizi, yeniNum)
		}
		fmt.Println(dizi)
		toplam := add(dizi...)
		fmt.Println("Dizideki sayıların toplamı:", toplam)
		ortalama := average(dizi...)
		fmt.Println("Dizideki sayıların ortalaması:", ortalama)
		enBuyuk := largest(dizi...)
		fmt.Println("Dizideki en büyük sayı:", enBuyuk)
		enKucuk := smallest(dizi...)
		fmt.Println("Dizideki en küçük sayı:", enKucuk)
	}
	func add(nums ...int) int {
		var toplam int = 0
		for i := 0; i < len(nums); i++ {
			toplam += nums[i]
		}
		return toplam
	}
	func average(nums ...int) int {
		var toplam int
		var toplam2 int
		for i := 0; i < len(nums); i++ {
			toplam += nums[i]
			toplam2 = toplam / len(nums)
		}
		return toplam2
	}

	func largest(nums ...int) int {

		largest := nums[0]

		for i := 0; i < len(nums); i++ {
			if nums[i] > largest {
				largest = nums[i]
			}
		}
		return largest
	}

	func smallest(nums ...int) int {

		smallest := nums[0]

		for i := 0; i < len(nums); i++ {
			if nums[i] < smallest {
				smallest = nums[i]
			}
		}
		return smallest
	}*/

	/*girdi := "Lorem Ipsum, basım ve dizgi endüstrisinin basit bir şekilde sahte metnidir."
	  	for index, element := range tekrarlama(girdi) {
	  		fmt.Println(index, "=", element)
	  	}

	  }

	  func tekrarlama(adet string) map[string]int {
	  	girdi := strings.Fields(adet)
	  	devam := make(map[string]int)
	  	for _, kelime := range girdi {
	  		_, esles := devam[kelime]
	  		if esles {
	  			devam[kelime] += 1
	  		} else {
	  			devam[kelime] = 1
	  		}
	  	}
	  	return devam
	  }*/

}
