package main
import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&m)
		arr := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&arr[j])
		}
		
		for j := 0; j < m-1; j++ {
			idx_min := j
			for k := j + 1; k < m; k++ {
				if arr[k] < arr[idx_min] {
					idx_min = k
				}
			}
			arr[j], arr[idx_min] = arr[idx_min], arr[j]
		}
		
		for j := 0; j < m; j++ {
			fmt.Print(arr[j], " ")
		}
		fmt.Println()
	}
}