package main

/*import "fmt"

type myString string

func main() {
	var ögrenciAdi myString = "Ahmet"
	ögrenciAdi.change("Mehmet")
	fmt.Println(ögrenciAdi)

}

func (m *myString) change(yeniVeri string) {

	if yeniVeri == "" {

	} else {
		*m = myString(yeniVeri)
	}
}*/

import "fmt"

func main() {
	str1 := "YazılımSeviyorum"
	charDict := make(map[rune]int)

	for _, c := range str1 {
		count := 0
		for _, ch := range str1 {
			if ch == c {
				count++
			}
		}
		charDict[c] = count
	}

	fmt.Println("Sonuc:")
	for key, value := range charDict {
		fmt.Printf("%c: %d\n", key, value)
	}
}
