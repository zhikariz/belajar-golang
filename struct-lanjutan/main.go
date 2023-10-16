package main

import "fmt"

type student struct {
	name  string
	grade int
}

func main() {
	student1 := student{
		name:  "Helmi",
		grade: 1,
	}

	student2 := &student1

	fmt.Println("student 1 name :", student1.name)
	fmt.Println("student 2 name :", student2.name)

	student2.name = "Adi"

	fmt.Println("student 1 name :", student1.name)
	fmt.Println("student 2 name :", student2.name)
}
