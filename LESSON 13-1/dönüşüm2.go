package main

import "fmt"

type myInt32 int32

func main() {
	var x myInt32 = 10
	var y myInt32 = 20

	var result myInt32 = add(x, y)
	fmt.Println(result)

	result2 := myInt32.myInt32add(x, y)
	fmt.Println(result2)

}

func (m myInt32) myInt32add(y myInt32) myInt32 {
	return m + y
}
func add(x myInt32, y myInt32) myInt32 {
	return x + y
}
