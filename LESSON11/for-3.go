package main

import "fmt"

func main() {
	notlar := make(map[int]float32)

	notlar[10] = 54
	notlar[11] = 34
	notlar[12] = 67

	for k, v := range notlar {
		fmt.Printf("%d-%.3f\n", k, v)
	}

	notlarS := []float32{54, 34, 67}
	for i, v := range notlarS {
		fmt.Printf("%d-%.3f\n", i, v)
	}

	var toplamNot float32 = 0

	for _, v := range notlar {
		toplamNot += v
	}
	fmt.Println(toplamNot)
}
