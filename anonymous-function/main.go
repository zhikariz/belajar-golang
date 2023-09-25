package main

import "fmt"

func main() {
	penjumlahan := func(x, y int) int {
		return x + y
	}

	hasil := penjumlahan(3, 4)
	fmt.Println(hasil) // Output: 7
}
