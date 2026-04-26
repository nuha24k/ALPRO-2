package main

import "fmt"

// Tipe bentukan array sesuai modul
type arrBalita [100]float64

// Subprogram untuk menghitung nilai minimum dan maksimum dengan pointer
func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

// Subprogram untuk menghitung rerata
func rerata(arrBerat arrBalita, n int) float64 {
	var sum float64
	for i := 0; i < n; i++ {
		sum += arrBerat[i]
	}
	return sum / float64(n)
}

func main() {
	var n int
	var arr arrBalita
	var bMin, bMax float64

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&arr[i])
	}

	// Pemanggilan subprogram
	hitungMinMax(arr, n, &bMin, &bMax)
	rata := rerata(arr, n)

	// Menampilkan output sesuai format yang diminta
	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}
