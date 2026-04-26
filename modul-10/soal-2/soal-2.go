package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	ikan := make([]float64, x)
	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	var totalBerat float64
	var jumlahWadah int

	// Menghitung berat per wadah
	if y > 0 {
		for i := 0; i < x; i += y {
			var beratWadah float64
			for j := 0; j < y && i+j < x; j++ {
				beratWadah += ikan[i+j]
			}
			fmt.Printf("%.2f ", beratWadah)
			totalBerat += beratWadah
			jumlahWadah++
		}
	}
	fmt.Println()

	// Menghitung rata-rata berat ikan di setiap wadah
	if jumlahWadah > 0 {
		rataRata := totalBerat / float64(jumlahWadah)
		fmt.Printf("%.2f\n", rataRata)
	} else {
		fmt.Printf("%.2f\n", 0.0)
	}
}
