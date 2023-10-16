package main

import "fmt"

type student struct {
	person
	grade int
}

func main() {
	var student1 student

	student1.name = "Helmi"
	student1.grade = 2
	student1.age = 17

	fmt.Println("name :", student1.name)
	fmt.Println("age :", student1.age)
	fmt.Println("grade :", student1.grade)
}
