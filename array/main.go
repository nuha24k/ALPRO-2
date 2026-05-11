package main

import "fmt"

func main() {
	arrLeft := [6]int{9, 10, 7, 11, 5, 8}
	// fmt.Println("Array:", arr)
	// fmt.Printf("Array: %v\n", arr)
	// fmt.Print(arr[0], " ", arr[1], " ", arr[2], " ", arr[3], " ", arr[4], "\n")

	// for i := 0; i < len(arr); i++ {
	// 	fmt.Printf("Index %d: %d\n", i, arr[i])
	// }

	arrRight := [6]int{7, 11, 8, 8, 5, 9}

	fmt.Println("Array Right:", arrRight)

	min := 0 

	if arrLeft[0] < arrRight[0] {
		min = arrLeft[0]
	} else {
		min = arrRight[0]
	}

	// fmt.Println("\nBest number of each index:")
	// for i := 0; i < len(arrLeft); i++ {
	// 	var best int
	// 	if arrLeft[i] > arrRight[i] {
	// 		best = arrLeft[i]
	// 	} else {
	// 		best = arrRight[i]
	// 	}
	// 	fmt.Printf("Index %d: %d\n", i, best)
	// }

	// for i := 0; i < len(arrRight); i++ {
	// 	fmt.Printf("Index %d: %d\n", i, arrRight[i])
	// }

}
