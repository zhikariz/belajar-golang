package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	student1 := struct {
		person
		grade int
	}{}

	student1.person = person{
		name: "Helmi",
		age:  17,
	}
	student1.grade = 2

	fmt.Println("name :", student1.person.name)
	fmt.Println("age :", student1.person.age)
	fmt.Println("grade :", student1.grade)
}
