package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n, inPizza int
	fmt.Print("Banyak Topping: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		x := rand.Float64()
		y := rand.Float64()

		// Titik pusat pizza (0.5, 0.5) dengan jejari r=0.5
		dx := x - 0.5
		dy := y - 0.5

		// Uji menggunakan ketidaksamaan area lingkaran
		if (dx*dx)+(dy*dy) <= 0.25 {
			inPizza++
		}
	}

	pi := 4.0 * float64(inPizza) / float64(n)

	fmt.Printf("Topping pada Pizza: %d\n", inPizza)
	fmt.Printf("PI : %.10f\n", pi)
}