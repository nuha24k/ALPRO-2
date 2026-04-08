package main

import "fmt"

func cetakFaktor(n int, pembagi int) {
	if pembagi > n {
		return
	}
	if n % pembagi == 0 {
		fmt.Print(pembagi, " ")
	}
	cetakFaktor(n, pembagi+1)
}

func main() {
	var n int
	fmt.Scan(&n)
	cetakFaktor(n, 1)
}