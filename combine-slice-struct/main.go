package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	students := []person{
		{name: "Helmi", age: 17},
		{name: "Dian", age: 18},
	}

	for _, student := range students {
		fmt.Println("name :", student.name, " age :", student.age)
	}
}
