package main

import "fmt"

func main() {
	var arr = [21]int{}
	totalVote := 0
	validVote := 0
	for {
		var suara int
		fmt.Scan(&suara)
		if suara == 0 {
			break
		}
		if suara > 0 && suara <= 20 {
			validVote++
			arr[suara]++
		}
		totalVote++
	}
	fmt.Println("Suara masuk: ", totalVote)
	fmt.Println("Suara sah: ", validVote)
	for i := 1; i <= 20; i++ {
		if arr[i] > 0 {
			fmt.Printf("%d : %d\n", i, arr[i])
		}
	}
}