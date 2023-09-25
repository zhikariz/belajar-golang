package main

import (
	"fmt"
	"time"
)

type Result struct {
	Data string
	Err  error
}

func main() {
	// Membuat channel untuk menerima hasil dari tugas asynchronous
	resultChan := make(chan Result)

	// Memulai goroutine untuk menjalankan tugas asynchronous
	go taskAsync(resultChan)

	// Melakukan tugas lainnya di sini
	for i := 0; i < 1; i++ {
		fmt.Println("Melakukan tugas lainnya:", i)
		time.Sleep(time.Second)
	}

	// Menerima hasil dari channel dan memanggil callback
	result := <-resultChan
	if result.Err != nil {
		fmt.Println("Error:", result.Err)
	} else {
		fmt.Println("Hasil tugas asynchronous:", result.Data)
	}

	fmt.Println("Selesai")
}

func taskAsync(resultChan chan<- Result) {
	// Simulasikan tugas asynchronous dengan hasil "Selesai" setelah 3 detik
	time.Sleep(3 * time.Second)
	result := Result{
		Data: "Selesai",
		Err:  nil,
	}
	resultChan <- result
}
