package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	berat := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&berat[i])
	}

	if n > 0 {
		min := berat[0]
		max := berat[0]

		for i := 1; i < n; i++ {
			if berat[i] < min {
				min = berat[i]
			}
			if berat[i] > max {
				max = berat[i]
			}
		}
		fmt.Printf("%f %f\n", min, max)
	}
}
