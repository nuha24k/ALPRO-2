package main
import "fmt"

func main() {
	var x int
	arr := []int{}
	
	for {
		fmt.Scan(&x)
		if x == -5313 {
			break
		} else if x != 0 {
			arr = append(arr, x)
		} else {
			for j := 1; j < len(arr); j++ {
				temp := arr[j]
				k := j
				for k > 0 && temp < arr[k-1] {
					arr[k] = arr[k-1]
					k--
				}
				arr[k] = temp
			}
			
			n := len(arr)
			if n > 0 {
				if n%2 != 0 {
					fmt.Println(arr[n/2])
				} else {
					fmt.Println((arr[n/2-1] + arr[n/2]) / 2)
				}
			}
		}
	}
}