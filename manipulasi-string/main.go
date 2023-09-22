package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {
	firstName := "Helmi"
	lastName := "Adi Prasetyo"
	alamat := "Surakarta,Jawa Tengah"
	email := "kartotoyinx@gmail.com"
	ukuranBaju := "XL"
	// ttl := "Surakarta, 30 Februari 1959"

	// fullName := firstName + lastName
	pemisahanAlamat := strings.Split(alamat, ",")
	kota := pemisahanAlamat[0]
	provinsi := pemisahanAlamat[1]
	fmt.Println("kotaku adalah ", kota)
	fmt.Println("provinsiku adalah ", provinsi)
	pemisahanEmail := strings.Split(email, "@")
	username := pemisahanEmail[0]
	fmt.Println("usernamenya adalah", username)

	// hasil := strings.Replace(ttl, "Surakarta", "Sukoharjo", -1)
	hasil := strings.Replace(ukuranBaju, "X", "XX", -1)
	fmt.Println("Ukuran bajuku yang baru adalah ", hasil)
	// fmt.Println("alamatku yang bener adalah ", hasil)

	fullName := fmt.Sprintf("%s %s\n", firstName, lastName)

	namaKopi := "Kopi Kenangan Mantan"

	adakah := strings.Contains(strings.ToLower(namaKopi), "mantan")

	if adakah {
		fmt.Println("moveon mas")
	} else {
		fmt.Println("ga ada aman")
	}

	encodedName := base64.StdEncoding.EncodeToString([]byte(fullName))
	fmt.Println(encodedName)
	decodedName, _ := base64.StdEncoding.DecodeString(encodedName)
	fmt.Println(string(decodedName))
}
