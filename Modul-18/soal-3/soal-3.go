package main

import (
	"fmt"
	"math/rand"
)

type Domino struct {
	sisi1   int
	sisi2   int
	nilai   int
	isBalak bool
}

type Dominoes struct {
	domino [28]Domino
	jumlah int
}

func buatDominoes() Dominoes {
	var dominoes Dominoes
	idx := 0
	for i := 0; i <= 6; i++ {
		for j := i; j <= 6; j++ {
			dominoes.domino[idx] = Domino{
				sisi1:   i,
				sisi2:   j,
				nilai:   i + j,
				isBalak: i == j,
			}
			idx++
		}
	}
	dominoes.jumlah = idx
	return dominoes
}

func kocokKartu(dominoes *Dominoes) {
	for i := dominoes.jumlah - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		dominoes.domino[i], dominoes.domino[j] = dominoes.domino[j], dominoes.domino[i]
	}
}

func ambilKartu(dominoes *Dominoes) Domino {
	if dominoes.jumlah == 0 {
		return Domino{}
	}
	dominoes.jumlah--
	return dominoes.domino[dominoes.jumlah]
}

func gambarKartu(D Domino, suit int) int {
	if D.sisi1 == suit {
		return D.sisi2
	}
	if D.sisi2 == suit {
		return D.sisi1
	}
	return -1
}

func nilaiKartu(domino Domino) int {
	return domino.nilai
}

func cocok(kartu1, kartu2 Domino) bool {
	return kartu1.sisi1 == kartu2.sisi1 || kartu1.sisi1 == kartu2.sisi2 || kartu1.sisi2 == kartu2.sisi1 || kartu1.sisi2 == kartu2.sisi2
}

func galiKartu(dominoes *Dominoes, kartuSekarang Domino) Domino {
	for dominoes.jumlah > 0 {
		kartu := ambilKartu(dominoes)
		if cocok(kartu, kartuSekarang) {
			return kartu
		}
	}
	return Domino{}
}

func sepasangKartu(kartu1, kartu2 Domino) bool {
	return nilaiKartu(kartu1)+nilaiKartu(kartu2) == 12
}

func mainGapleh(dominoes *Dominoes) {
	fmt.Println("\n=== SIMULASI GAPLEH ===")
	if dominoes.jumlah == 0 {
		return
	}
	kartuSekarang := ambilKartu(dominoes)
	fmt.Printf("Awal: (%d,%d)\n", kartuSekarang.sisi1, kartuSekarang.sisi2)
	for dominoes.jumlah > 0 {
		kartuBaru := galiKartu(dominoes, kartuSekarang)
		if kartuBaru == (Domino{}) {
			fmt.Println("Tidak ada kartu yang cocok")
			break
		}
		fmt.Printf(
			"-> (%d,%d)\n",
			kartuBaru.sisi1,
			kartuBaru.sisi2,
		)
		kartuSekarang = kartuBaru
	}
}

func main() {
	dominoes := buatDominoes()
	kocokKartu(&dominoes)
	kartu := ambilKartu(&dominoes)
	fmt.Printf("Kartu yang diambil: (%d,%d)\n", kartu.sisi1, kartu.sisi2)
	fmt.Println("Nilai kartu =", nilaiKartu(kartu))
	fmt.Println("Balak =", kartu.isBalak)
	fmt.Printf("Jika suit = %d, nilainya = %d\n", kartu.sisi1, gambarKartu(kartu, kartu.sisi1))
	kartuCocok := galiKartu(&dominoes, kartu)
	fmt.Printf("Kartu hasil gali: (%d,%d)\n", kartuCocok.sisi1, kartuCocok.sisi2)
	fmt.Println("Apakah total nilainya 12?", sepasangKartu(kartu, kartuCocok))
	dominoGapleh := buatDominoes()
	kocokKartu(&dominoGapleh)
	mainGapleh(&dominoGapleh)
}