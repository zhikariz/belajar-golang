package main

import "fmt"

// first class function
func main() {
	var jml func(int, int) int
	jml = penjumlahan

	hasil := jml(2, 7)
	fmt.Println(hasil)

	jml = pengurangan

	hasilPengurangan := jml(5, 2)
	fmt.Println(hasilPengurangan)
}

func penjumlahan(bil1, bil2 int) int {
	return bil1 + bil2
}

func pengurangan(bil1, bil2 int) int {
	return bil1 - bil2
}
