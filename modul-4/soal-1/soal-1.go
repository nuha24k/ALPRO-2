package main

import "fmt"

func faktorial(n int, hasil *int) {
	*hasil = 1
	for i := 2; i <= n; i++ {
		*hasil *= i
	}
}

func permutasi(n, r int, hasil *int) {
	var nFaktorial, nMinRFaktorial int
	
	faktorial(n, &nFaktorial)
	faktorial(n-r, &nMinRFaktorial)

	*hasil = nFaktorial / nMinRFaktorial
}

func kombinasi(n, r int, hasil *int) {
	var nFaktorial, rFaktorial, nMinRFaktorial int

	faktorial(n, &nFaktorial)
	faktorial(r, &rFaktorial)
	faktorial(n-r, &nMinRFaktorial)
	
	*hasil = nFaktorial / (rFaktorial * nMinRFaktorial)
}

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)

	var pA, cA int
	permutasi(a, c, &pA)
	kombinasi(a, c, &cA)
	
	var pB, cB int
	permutasi(b, d, &pB)
	kombinasi(b, d, &cB)

	fmt.Println(pA, cA)
	
	fmt.Println(pB, cB)
}
