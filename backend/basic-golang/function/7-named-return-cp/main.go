package main

import "fmt"

//fungsi square sama seperti sebelumnya
//yang membedakan adalah fungsi ini
//menggunakan named return
func main() {
	result1, result2 := square(4, 5)
	fmt.Printf("%d dan %d\n", result1, result2)
	fmt.Println(square(9, 8))
}

func square(angka1, angka2 int) (result1, result2 int) {
	result1 = angka1 * angka1
	result2 = angka2 * angka2
	return result1, result2
}
