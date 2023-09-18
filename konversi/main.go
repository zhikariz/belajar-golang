package main

import (
	"fmt"
	"strconv"
)

func main() {
	// string ke integer
	id := "1094"
	idInteger, _ := strconv.Atoi(id)
	fmt.Printf("%T\n", idInteger)
	// integer ke string
	uang := 6000
	stringUang := strconv.Itoa(uang)
	stringUang2 := fmt.Sprintf("%d", uang)
	fmt.Printf("%T\n", stringUang)
	fmt.Printf("%T\n", stringUang2)
	// float ke string
	nilaiFloat := 3.14
	stringFloat := strconv.FormatFloat(nilaiFloat, 'f', -1, 64)
	stringFloat2 := fmt.Sprintf("%v", nilaiFloat)
	fmt.Printf("valuenya adalah %v tipe datanya adalah %T\n", stringFloat, stringFloat)
	fmt.Printf("valuenya adalah %v tipe datanya adalah %T\n", stringFloat2, stringFloat2)
	// string ke float
	stringFloat3 := "6.12"
	nilaiFloat3, _ := strconv.ParseFloat(stringFloat3, 32)
	fmt.Printf("valuenya adalah %v tipe datanya adalah %T\n", float32(nilaiFloat3), float32(nilaiFloat3))
}
