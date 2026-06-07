package main
import "fmt"

func main() {
	var dat, sum float64
	var count int

	fmt.Scan(&dat)
	for dat != 9999 {
		sum += dat
		count++
		fmt.Scan(&dat)
	}

	if count > 0 {
		fmt.Printf("Rerata: %.2f\n", sum/float64(count))
	} else {
		fmt.Println("Tidak ada data")
	}
}
