package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x, hapusIdx, cari int
	fmt.Print("Masukkan jumlah elemen N: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("a. Keseluruhan isi array:", arr)

	fmt.Print("b. Indeks ganjil: ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("c. Indeks genap: ")
	for i := 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("Masukkan angka x untuk kelipatan: ")
	fmt.Scan(&x)
	fmt.Print("d. Indeks kelipatan ", x, ": ")
	for i := 0; i < len(arr); i++ {
		if x != 0 && i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&hapusIdx)
	if hapusIdx >= 0 && hapusIdx < len(arr) {
		arr = append(arr[:hapusIdx], arr[hapusIdx+1:]...) // Menghapus menggunakan fitur slice
	}
	fmt.Println("e. Array setelah dihapus:", arr)

	sum := 0.0
	for _, val := range arr {
		sum += float64(val)
	}
	rata := sum / float64(len(arr))
	fmt.Printf("f. Rata-rata: %.2f\n", rata)

	var varSum float64
	for _, val := range arr {
		varSum += math.Pow(float64(val)-rata, 2)
	}
	stdDev := math.Sqrt(varSum / float64(len(arr)))
	fmt.Printf("g. Standar Deviasi: %.2f\n", stdDev)

	fmt.Print("Masukkan bilangan untuk dicari frekuensinya: ")
	fmt.Scan(&cari)
	freq := 0
	for _, val := range arr {
		if val == cari {
			freq++
		}
	}
	fmt.Printf("h. Frekuensi nilai %d: %d\n", cari, freq)
}