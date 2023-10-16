package main

import "fmt"

func main() {
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("nilai", numberA)
	fmt.Println("alamat memory", &numberA)

	fmt.Println("nilai", *numberB)
	fmt.Println("alamat memory", numberB)

	// numberA = 8
	ganti(numberB, 9.6)

	fmt.Println("nilai", numberA)
	fmt.Println("alamat memory", &numberA)

	fmt.Println("nilai", *numberB)
	fmt.Println("alamat memory", numberB)
}

func ganti(nilaiAwal *int, nilai float64) {
	*nilaiAwal = int(nilai)
}
