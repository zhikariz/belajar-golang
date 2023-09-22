package main

import "fmt"

func main() {
	var hari = "Rabu"

	switch hari {
	case "Senin":
		fmt.Println("Hari ini adalah Senin")
	case "Selasa":
		fmt.Println("Hari ini adalah Selasa")
	case "Rabu":
		fmt.Println("Hari ini adalah Rabu")
	case "Kamis":
		fmt.Println("Hari ini adalah Kamis")
	case "Jumat":
		fmt.Println("Hari ini adalah Jumat")
	case "Sabtu":
		fmt.Println("Hari ini adalah Sabtu")
	case "Minggu":
		fmt.Println("Hari ini adalah Minggu")
	default:
		fmt.Println("Hari ini bukan hari dalam seminggu")
	}
}
