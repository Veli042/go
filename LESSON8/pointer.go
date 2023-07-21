package main

import "fmt"

/*func main() {
	//pointersiz hali
	var örneksayi int
	örneksayi = 85
	fmt.Printf("Dışarıdaki sayı ilk durum: %d\n", örneksayi)
	artir(örneksayi)
	fmt.Printf("Dışarıdaki sayı son durum: %d\n", örneksayi)
}

func artir(sayi1 int) int{
	fmt.Printf("İçereideki sayı ilk durum: %d\n", sayi1)
	sayi1 = sayi1 + 1
	fmt.Printf("İçereideki sayı son durum: %d\n", sayi1)
	return sayi1


}*/

func main() {
	var örneksayi int
	örneksayi = 100
	örneksayiP := &örneksayi
	fmt.Printf("Dışarıdaki sayının adresi: %p\n", örneksayiP)
	fmt.Printf("Dışarıdaki sayı ilk durum: %d\n", örneksayi)
	artir(örneksayiP)
	fmt.Printf("Dışarıdaki sayı son durum: %d\n", örneksayi)

}

func artir(sayi1P *int) {
	fmt.Printf("İçerideki sayının adresi: %p\n", sayi1P)
	fmt.Printf("İçereideki sayı ilk durum: %d\n", *sayi1P)
	*sayi1P = *sayi1P + 1
	fmt.Printf("İçereideki sayı son durum: %d\n", *sayi1P)
}
