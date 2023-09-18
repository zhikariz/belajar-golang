package main

import (
	"fmt"
	"strconv"
)

func main() {
	// bikin 3 inputan : bilangan 1, bilangan 2, operasi / aksi
	// pencabangan apakah operasi / aksi berupa penjumlahan/pembagian/pengurangan/sisa hasil (modulus)
	// tampilkan hasilnya
	var inputBilangan1, inputBilangan2, inputOperasi string

	fmt.Println("Silakan masukkan bilangan 1: ")
	fmt.Scanln(&inputBilangan1)

	bilangan1, err := strconv.Atoi(inputBilangan1)
	if err != nil {
		fmt.Println("masukkan hanya integer / angka saja")
		return
	}

	fmt.Println("Silakan masukkan bilangan 2: ")
	fmt.Scanln(&inputBilangan2)

	bilangan2, err := strconv.Atoi(inputBilangan2)
	if err != nil {
		fmt.Println("masukkan hanya integer / angka saja")
		return
	}

	fmt.Println(`
	Operasi / Aksi :
	1. Penjumlahan
	2. Pengurangan
	3. Pembagian
	4. Perkalian
	5. Modulus
	`)

	fmt.Println("Silakan masukkan aksi (1-5)")
	fmt.Scanln(&inputOperasi)

	operasi, err := strconv.Atoi(inputOperasi)

	if err != nil {
		fmt.Println("masukkan hanya integer / angka saja")
		return
	}

	var hasil int

	if operasi == 1 {
		hasil = bilangan1 + bilangan2
	} else if operasi == 2 {
		hasil = bilangan1 - bilangan2
	} else if operasi == 3 {
		hasil = bilangan1 / bilangan2
	} else if operasi == 4 {
		hasil = bilangan1 * bilangan2
	} else if operasi == 5 {
		hasil = bilangan1 % bilangan2
	}

	fmt.Println("Bilangan 1 : ", bilangan1)
	fmt.Println("Bilangan 2 : ", bilangan2)
	fmt.Println("Aksi/Operasi  : ", inputOperasi)
	fmt.Println("Hasil dari Operasi tsb adalah : ", hasil)
}
