package main

import "fmt"

// Check Point:
// Search Roles
// - Input: Role
// - Output: Rearch Result

// Contoh:
// Data Users: [{Aditira 20 Programmer} {Dimas 18 Designer} {Yuni 21 DevOps} {Dito 22 Programmer} {Juno 25 DevOps}]

// Input:
//   - Masukan Role : Programmer

// Output:
// Programmer Found:
// Name:  Aditira  Age:  20  Role:  Programmer
// Name:  Dito  Age:  22  Role:  Programmer

// Input:
//   - Masukan Role : Secretary

// Output:
// Role: Sercretary Not Found!

type User struct {
	name string
	age  int
	role string
}

func main() {

	users := []User{
		{
			name: "Aditira",
			age:  20,
			role: "Programmer",
		},
		{
			name: "Dimas",
			age:  18,
			role: "Designer",
		},
		{
			name: "Yuni",
			age:  21,
			role: "DevOps",
		},
		{
			name: "Dito",
			age:  22,
			role: "Programmer",
		},
		{
			name: "Juno",
			age:  25,
			role: "DevOps",
		},
	}

	// TODO: answer here
	var roles string
	fmt.Scan(&roles)
	userFound := []User{}

	for _, v := range users {
		if v.role == roles {
			var userStruct User = User{
				name: v.name,
				age:  v.age,
				role: v.role,
			}
			userFound = append(userFound, userStruct)

		}

	}
	if len(userFound) == 0 {
		fmt.Println("Role not found")
	} else {
		for _, v := range userFound {
			fmt.Println("name: ", v.name, "age: ", v.age, "role: ", v.role)
		}
	}

}
