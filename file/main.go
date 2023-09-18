package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("file/halo.txt")
	if err != nil {
		fmt.Println("errorsnya adalah ", err)
		return
	}
	n, err := file.WriteString("Halo namaku helmi, aku sedang belajar golang fundamental")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("n adalah ", n)
	err = file.Sync()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("filenya berhasil dibuat loh ", file.Name())
}
