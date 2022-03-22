package main

import "fmt"

//fungsi printWord akan melakukan print satu persatu nilai parameter yang diterimanya
//contoh nilai parameter yang diterima
//("selamat","pagi","siang",sore)
//outputnya
//selamat
//pagi
//siang
//sore
func main() {
	printWord("halo")
	printWord("halo", "selamat siang")
	printWord("halo", "andi", "dan", "bagus")
	printWord("mencoba", "variadic", "param", "pada", "go")
}

func printWord(words ...string) {
	result := ""
	for _, words := range words {
		result += words + "\n"
	}
	fmt.Println(result)
}
