package main

import "fmt"

func main() {
	var arr = [21]int{}
	totalVote := 0
	validVote := 0
	ketua := 0
	wakilKetua := 0
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
	for i := 1; i <= validVote; i++ {
		if arr[i] > arr[ketua] {
			ketua = i
		}
		if arr[i] > arr[wakilKetua] && i != ketua {
			wakilKetua = i
		}
	}
	fmt.Println("Ketua RT: ", ketua)
	fmt.Println("Wakil Ketua: ", wakilKetua)
}