//dikerjakan jika waktu cukup saja
package main

import "fmt"

type Employee struct {
	Name              string
	BaseSalary, Bonus int
}

func (e Employee) GetFullSalary() int {
	return e.BaseSalary + e.Bonus
}
func (e *Employee) AddBonus(bonus int) {
	e.Bonus += bonus
}

//method GetFullSalary() untuk menampilkan salary + bonus
//method AddBonus(bonus int) untuk mengubah nilai atribut bonus employee

// TODO: answer here

func main() {
	employee1 := Employee{Name: "bob", BaseSalary: 4000000, Bonus: 300000}
	fmt.Println("sebelum bonus dinaikkan")
	fmt.Println(employee1.GetFullSalary())
	fmt.Println("tambah 100000 ke bonus")
	employee1.AddBonus(100000)
	fmt.Println("setelah bonus dinaikkan")
	fmt.Println(employee1.GetFullSalary())
}
