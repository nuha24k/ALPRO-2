package main

import (
	"fmt"
)

type MesinKarakter struct {
	data string
	pos  int
}

func start(mesinKarakter *MesinKarakter, teks string) {
	mesinKarakter.data = teks
	mesinKarakter.pos = 0
}

func maju(mesinKarakter *MesinKarakter) {
	if mesinKarakter.pos < len(mesinKarakter.data) {
		mesinKarakter.pos++
	}
}

func eop(mesinKarakter MesinKarakter) bool {
	if mesinKarakter.pos >= len(mesinKarakter.data) {
		return true
	}
	return mesinKarakter.data[mesinKarakter.pos] == '.'
}

func cc(mesinKarakter MesinKarakter) byte {
	return mesinKarakter.data[mesinKarakter.pos]
}

func bacaSemua(teks string) {
	var mesinKarakter MesinKarakter
	start(&mesinKarakter, teks)
	for !eop(mesinKarakter) {
		fmt.Printf("%c", cc(mesinKarakter))
		maju(&mesinKarakter)
	}
	fmt.Println()
}

func hitungKarakter(teks string) int {
	var mesinKarakter MesinKarakter
	start(&mesinKarakter, teks)
	jumlah := 0
	for !eop(mesinKarakter) {
		jumlah++
		maju(&mesinKarakter)
	}
	return jumlah
}

func hitungA(teks string) int {
	var mesinKarakter MesinKarakter
	start(&mesinKarakter, teks)
	jumlah := 0
	for !eop(mesinKarakter) {
		if cc(mesinKarakter) == 'A' {
			jumlah++
		}
		maju(&mesinKarakter)
	}
	return jumlah
}

func frekuensiA(teks string) float64 {
	total := hitungKarakter(teks)
	if total == 0 {
		return 0
	}
	return float64(hitungA(teks)) / float64(total)
}

func hitungLE(teks string) int {
	jumlah := 0
	for i := 0; i < len(teks)-1; i++ {
		if teks[i] == '.' {
			break
		}
		if teks[i] == 'L' &&
			teks[i+1] == 'E' {
			jumlah++
		}
	}
	return jumlah
}

func main() {
	teks := "SAYA LEPAS LELE DAN LEBAH."
	fmt.Print("Karakter terbaca: ")
	bacaSemua(teks)
	fmt.Println("Jumlah karakter =", hitungKarakter(teks))
	fmt.Println("Jumlah huruf A =", hitungA(teks))
	fmt.Println("Frekuensi A =", frekuensiA(teks))
	fmt.Println("Jumlah kata LE =", hitungLE(teks))
}