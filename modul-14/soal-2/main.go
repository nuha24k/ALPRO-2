package main
import "fmt"

func main() {
	var n, m, x int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&m)
		ganjil := []int{}
		genap := []int{}
		for j := 0; j < m; j++ {
			fmt.Scan(&x)
			if x%2 != 0 {
				ganjil = append(ganjil, x)
			} else {
				genap = append(genap, x)
			}
		}
		
		for j := 0; j < len(ganjil)-1; j++ {
			idx := j
			for k := j + 1; k < len(ganjil); k++ {
				if ganjil[k] < ganjil[idx] { idx = k }
			}
			ganjil[j], ganjil[idx] = ganjil[idx], ganjil[j]
		}
		
		for j := 0; j < len(genap)-1; j++ {
			idx := j
			for k := j + 1; k < len(genap); k++ {
				if genap[k] > genap[idx] { idx = k }
			}
			genap[j], genap[idx] = genap[idx], genap[j]
		}
		
		for _, v := range ganjil { fmt.Print(v, " ") }
		for _, v := range genap { fmt.Print(v, " ") }
		fmt.Println()
	}
}