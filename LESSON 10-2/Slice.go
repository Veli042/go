package main

import "fmt"

func main() {

	// syntax , tanımlama şekli
	/*var numaralar []int = []int{}
	var sayilar []int = []int{2, 3, 4}

	ögrenciler := make([]string, 4)*/

	/*var num []int = []int{2, 3, 4, 6, 7}
	fmt.Println(num)

	num = append(num, 5)

	fmt.Println("güncel hali", num)*/

	// Slice ile toplama işlemi

	var num []int = []int{}
	var yeniNum int
	for {
		fmt.Println("Sayı giriniz:")
		fmt.Scan(&yeniNum)
		if yeniNum == 99 {
			break
		}
		num = append(num, yeniNum)
	}
	fmt.Println(num)
	toplam := add(num...)
	fmt.Println(toplam)

}

func add(nums ...int) int {
	var toplam int = 0
	for i := 0; i < len(nums); i++ {
		toplam += nums[i]
	}
	return toplam
}
