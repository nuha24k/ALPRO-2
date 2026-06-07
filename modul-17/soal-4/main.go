package main
import (
	"fmt"
	"math"
)

func main() {
	var pi, piLama, suku float64
	var i int = 1
	var limit int

	fmt.Print("N suku pertama: ")
	fmt.Scan(&limit)

	for {
		if i%2 != 0 {
			suku = 1.0 / float64(2*i-1)
		} else {
			suku = -1.0 / float64(2*i-1)
		}

		piLama = pi
		pi += suku
		estimasiSekarang := pi * 4
		estimasiLama := piLama * 4

		if i > 1 && math.Abs(estimasiSekarang-estimasiLama) <= 0.00001 {
			fmt.Printf("Hasil PI: %.10f\n", estimasiSekarang)
			fmt.Printf("Pada i ke: %d\n", i)
			break
		}
		i++
	}
}