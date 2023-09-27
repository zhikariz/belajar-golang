package main

import "fmt"

func main() {
	var fruits = []string{"apple", "grape", "banana", "melon"}
	var newFruits = fruits[0:2]

	fmt.Println(newFruits)
}
