package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	defer catch()
	var name string
	fmt.Print("Type your name:")
	fmt.Scanln(&name)

	err := validate(name)
	if err != nil {
		panic(err.Error())
		fmt.Println("end")
	}
	fmt.Println("halo", name)
}

func validate(input string) error {
	if strings.TrimSpace(input) == "" {
		return errors.New("input cannot be empty")
	}
	return nil
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application Running Perfectly")
	}
}
