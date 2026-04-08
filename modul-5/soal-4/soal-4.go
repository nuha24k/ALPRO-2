package main

import "fmt"

func cetakPolaCermin(n int) {
	if n == 1 {
		fmt.Print(1, " ")
	} else {
		fmt.Print(n, " ")
		cetakPolaCermin(n - 1)
		fmt.Print(n, " ")
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	cetakPolaCermin(n)
}