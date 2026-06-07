package main
import "fmt"

const nMax int = 7919
type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}
type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, 
                 &pustaka[i].penerbit, &pustaka[i].eksemplar, &pustaka[i].tahun, 
                 &pustaka[i].rating)
	}
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		temp := pustaka[i]
		j := i
		for j > 0 && temp.rating > pustaka[j-1].rating {
			pustaka[j] = pustaka[j-1]
			j--
		}
		pustaka[j] = temp
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	fmt.Println(pustaka[0].judul, pustaka[0].penulis, pustaka[0].penerbit, pustaka[0].tahun)
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	limit := 5
	if n < 5 { limit = n }
	for i := 0; i < limit; i++ {
		fmt.Println(pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	kiri, kanan := 0, n-1
	ketemu := false
	for kiri <= kanan && !ketemu {
		tengah := (kiri + kanan) / 2
		if pustaka[tengah].rating == r {
			fmt.Println(pustaka[tengah].judul, pustaka[tengah].penulis, 
                        pustaka[tengah].penerbit, pustaka[tengah].tahun, 
                        pustaka[tengah].eksemplar, pustaka[tengah].rating)
			ketemu = true
		} else if r > pustaka[tengah].rating {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}
	if !ketemu { fmt.Println("Tidak ada buku dengan rating seperti itu") }
}