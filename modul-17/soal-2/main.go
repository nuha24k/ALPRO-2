package main
import "fmt"

func main() {
	var x, str string
	var n, count int
	var pos []int

	fmt.Scan(&x, &n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&str)
		if str == x {
			count++
			pos = append(pos, i)
		}
	}

	ada := count > 0
	fmt.Println("a. Apakah string x ada?", ada)
	
	fmt.Print("b. Posisi: ")
	if ada {
		for _, p := range pos {
			fmt.Print(p, " ")
		}
		fmt.Println()
	} else {
		fmt.Println("-")
	}
	
	fmt.Println("c. Jumlah string x:", count)
	fmt.Println("d. Adakah sedikitnya dua?", count >= 2)
}
