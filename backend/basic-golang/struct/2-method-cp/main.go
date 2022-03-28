package main

import "fmt"

// buat struct Rectangle (persegi panjang) dengan dua atribut yaitu Width dan Length
// tambah dua method :
// GetArea() dan GetPerimeter()
// GetArea() digunakan untuk menampilkan (print) luas persegi panjang
// GetPerimeter() digunakan untuk menampilkan (print) keliling persegi panjang

type Rectangle struct {
	Width  int
	Length int
}

func (r Rectangle) GetArea() {
	area := r.Width * r.Length
	fmt.Printf("%d \n", area)
}
func (r Rectangle) GetPerimeter() {
	area := (2 * r.Width) + (2 * r.Length)
	fmt.Printf("%d", area)
}

// TODO: answer here
func main() {
	var r Rectangle
	r.Width = 10
	r.Length = 20
	fmt.Println(r)

	r.GetArea()
	r.GetPerimeter()
}
