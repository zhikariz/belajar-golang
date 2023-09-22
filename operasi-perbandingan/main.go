package main

import "fmt"

func main() {
	tbA := 140
	tbB := 140

	if tbA > tbB {
		fmt.Println("A lebih tinggi dari B")
	} else {
		fmt.Println("A lebih kecil dari B")
	}

	nilai := 60
	lulus := false

	if nilai >= 100 {
		lulus = true
		fmt.Println("Cumlaude")
	} else if nilai >= 80 {
		lulus = true
		fmt.Println("Bagus")
	} else if nilai >= 60 {
		lulus = true
		fmt.Println("Cukup")
	} else {
		lulus = false
		fmt.Println("Belajar Lagi")
	}

	if tbA > 150 || tbB > 150 {
		fmt.Println("kalian boleh ikut basket")
	} else {
		fmt.Println("kalian ga boleh ikut basket")
	}

	if !lulus {
		fmt.Println("dimarahin mama")
	} else {
		fmt.Println("dikasih uang")
	}
}
