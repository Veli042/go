package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for {
		r1 := rand.Intn(100)

		fmt.Println(r1)
		if r1 > 76 {
			break
		}
	}
	fmt.Println("İşlem Bitti.")

}

func farkli() {

	var r2 int = 0

	for r2 < 90 {
		r2 = rand.Intn(100)
		fmt.Println(r2)
	}
}
