package main
import "fmt"

func main() {
	var x int
	arr := []int{}
	
	for {
		fmt.Scan(&x)
		if x < 0 { break }
		arr = append(arr, x)
	}
	
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		j := i
		for j > 0 && temp < arr[j-1] {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = temp
	}
	
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()
	
	isKonsisten := true
	jarak := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != jarak {
			isKonsisten = false
			break
		}
	}
	
	if isKonsisten {
		fmt.Printf("Data berjarak %d\n", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}