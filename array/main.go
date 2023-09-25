package main

import "fmt"

func main() {
	names := [4]string{"trafalgar", "d", "water", "law"}

	fmt.Println(names[0], names[1], names[2], names[3])
	fmt.Println("jumlah array names adalah ", len(names))

	nama := "Helmi Adi Prasetyo"
	fmt.Println("jumlahnya adalah", string(nama[0]))
}
