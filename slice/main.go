package main

import "fmt"

func main() {
	var fruits = []string{"apple", "grape", "banana", "melon"}
	var newFruits = fruits[0:2]
	fmt.Println("ini data fruits pertama ", fruits)

	fruits2 := fruits[:]
	fmt.Println("ini data fruits kedua", fruits2)

	fruits[2] = "dragon fruit"
	fruits2[3] = "strawberry"
	fmt.Println("ini data fruits setelah di ganti", fruits)
	fmt.Println("ini data fruits2 setelah di ganti", fruits2)

	fmt.Println(newFruits)
}
