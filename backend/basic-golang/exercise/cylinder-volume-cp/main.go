package main

import (
	"fmt"
	"math"
)

// Check Point:
// Menghitung volume tabung
// - Input: jari-jari, tinggi
// - Output: volume tabung

// Contoh:
// Input:
// - Masukkan jari-jari alas tabung: 2
// - Masukkan tinggi tabung : 20
// Output:
// - Jadi volumenya adalah : 251.200012

func main() {

	var jari, tinggi, volume float32

	fmt.Print("Masukkan jari-jari alas tabung : ")
	fmt.Scanln(&jari)

	fmt.Print("Masukkan tinggi tabung : ")
	fmt.Scanln(&tinggi)

	volume = math.Pi * jari * jari * tinggi

	fmt.Println("Jadi volumenya adalah : ", volume)
}
