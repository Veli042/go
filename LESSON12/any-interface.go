package main

import "fmt"

func main() {

	var x interface{}
	x = 42
	x = "Konya"
	fmt.Println(x)

	var y any

	y = 42
	y = "Konya"
	y = 4 + 3i
	fmt.Println(y)
}
