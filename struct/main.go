package main

import "fmt"

// challenge bikin array of struct dari semua peserta bootcamp

func main() {
	type Peserta struct {
		NIM         string
		NamaLengkap string
		NoHp        int
		IPK         float64
		Hadir       bool
	}

	helmi := Peserta{
		NIM:         "A001",
		NamaLengkap: "Helmi Adi Prasetyo",
		NoHp:        812345678,
		IPK:         3.89,
		Hadir:       true,
	}

	fmt.Printf(`
	Saya adalah seorang mahasiswa yang bernama %s
	memiliki Nim %s berikut adalah no hp yang bisa dihubungi %d
	saya memiliki nilai ipk %g dan kehadiran saya adalah %t
	`, helmi.NamaLengkap, helmi.NIM, helmi.NoHp, helmi.IPK, helmi.Hadir)
}
