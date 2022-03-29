package main

import "fmt"

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan address operator dan indirect operator.

// Print alamat memori dari masing-masing variabel dibawah ini
func main() {
	name := "Roger"
	age := 20
	isMarried := true

	nameAddress := &name
	ageAddress := &age
	isMarriedAddress := &isMarried

	fmt.Printf("%v \n", nameAddress)
	fmt.Printf("%v \n", ageAddress)
	fmt.Printf("%v \n", isMarriedAddress)

	fmt.Printf("%s \n", *nameAddress)
	fmt.Printf("%d \n", *ageAddress)
	fmt.Printf("%t", *isMarriedAddress)
	// TODO: answer here
}
