package main

import "fmt"

func main() {
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
	toplam := topla(num...)
	fmt.Println(toplam)
}

func topla(nums ...int) int {
	var toplam int = 0

	for i := 0; i < len(nums); i++ {
		toplam += nums[i]
	}
	return toplam
}
