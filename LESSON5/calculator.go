package main

import (
	"errors"
	"fmt"
)

func main() {
	var (
		sayi1     float32
		sayi2     float32
		operation string
		result    float32
		err       error
	)
	for {
		fmt.Println("1. Sayıyı giriniz:")
		fmt.Scan(&sayi1)

		fmt.Println("2. Sayıyı giriniz:")
		fmt.Scan(&sayi2)

		fmt.Println("Yapmak istediğiniz işlemi yazınız(+,-,*,/):")
		fmt.Scan(&operation)
		if operation == "+" {
			result = add(sayi1, sayi2)
			fmt.Println("Sonuç:", result)
		} else if operation == "-" {
			result = extraction(sayi1, sayi2)
			fmt.Println("Sonuç:", result)
		} else if operation == "*" {
			result = impact(sayi1, sayi2)
			fmt.Println("Sonuç:", result)
		} else if operation == "/" {
			result, err = divide(sayi1, sayi2)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Sonuç:", result)
			}

		} else {
			fmt.Println("Hatalı işlem.")
		}
		fmt.Println("Devam etmek için (evet.hayır) yazınız:")
		fmt.Scan(&operation)

		if operation == "evet" {
			continue
		} else {
			break
		}
	}
	fmt.Println("İşleminiz sonlandırılıyor..")
	fmt.Println("İşlem Bitti :)")

}

func add(a float32, b float32) float32 {
	return a + b
}

func extraction(a float32, b float32) float32 {
	return a - b
}

func impact(a float32, b float32) float32 {
	return a * b
}

func divide(a float32, b float32) (bolum float32, err error) {

	if b == 0 {
		err = errors.New("Payda 0 olamaz..")
	}
	bolum = a / b

	return bolum, err

}
