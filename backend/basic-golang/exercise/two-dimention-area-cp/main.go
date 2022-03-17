package main

import (
	"fmt"
	"math"
)

// Check Point:
// Two Dimention Area
// - Input: Select Choice
// - Output: Result Operation

// Contoh:
// Input:
// 1: Rectange Area
// 2: Rectangular Area
// 3: Triangle Area
// 4: Circle Area
// - Enter choice 1, 2, 3, or 4: 1 | 2 | 3 | 4
//   - (1) Rectange Area
//   	- Masukkan sisi : 5
//   - (2) Rectangular Area
// 		- Masukkan panjang : 5
// 		- Masukkan lebar : 10
// 	 - (3) Triangle Area
// 		- Masukkan panjang alas segitiga: 5
// 		- Masukkan tinggi segitiga: 10
// 	 - (4) Circle Area
//      - Masukkan jari-jari : 4

// Output:
// - (1) Luas Persegi adalah : 25
// - (2) Luas Persegi Panjang adalah : 50
// - (3) Luas Segitiga adalah : 25
// - (4) Luas Lingkaran adalah : 50.240000

func main() {
	var num int
	var sisi float64 = 0
	var sisi2 float64 = 0
	fmt.Println("Input:")
	fmt.Print("1: Rectange Area \n2: Rectangular Area \n3: Triangle Area \n4: Circle Area \n")
	fmt.Println("Enter choice 1, 2, 3, or 4: 1 | 2 | 3 | 4")
	fmt.Scanln(&num)
	switch num {
	case 1:
		fmt.Println("Masukkan sisi")
		fmt.Scan(&sisi)
		var area = sisi * sisi
		fmt.Print(area)
	case 2:
		fmt.Println("Masukkan alas dan panjang")
		fmt.Scan(&sisi)
		fmt.Scan(&sisi2)
		var area = sisi * sisi2
		fmt.Print(area)
	case 3:
		fmt.Println("Masukkan alas dan tinggi")
		fmt.Scan(&sisi)
		fmt.Scan(&sisi2)
		var area = (sisi * sisi2) / 2
		fmt.Print(area)
	case 4:
		fmt.Println("Masukkan jari jari")
		fmt.Scan(&sisi)
		fmt.Print(math.Pi * sisi * sisi)
	}
}
