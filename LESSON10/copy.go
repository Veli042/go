package main

import "fmt"

func main() {

	sayilar := []int{3, 5, 4}
	nums := make([]int, len(sayilar))
	copy(nums, sayilar)
	fmt.Println(nums)
}
