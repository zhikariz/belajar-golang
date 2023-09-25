package main

import "fmt"

func apply(hasil func(int) int, bilangan int) int {
	return hasil(bilangan)
}

func double(bilangan int) int {
	return bilangan * 2
}

func main() {
	result := apply(double, 5)
	fmt.Println(result) // Output: 10
}
