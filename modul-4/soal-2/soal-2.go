package main

import "fmt"

func hitungSkor(soal *int, skor *int) {
	*soal = 0
	*skor = 0
	var waktu int
	
	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)
		if waktu < 301 {
			*soal++         
			*skor += waktu   
		}
	}
}

func main() {
	var namaPemenang string 
	var maxSoal int = -1    
	var minSkor int = 99999 

	var nama string
	for {
		fmt.Scan(&nama)
		
		if nama == "Selesai" {
			break
		}

		var soal, skor int
		
		hitungSkor(&soal, &skor)

		if soal > maxSoal {
			maxSoal = soal
			minSkor = skor
			namaPemenang = nama
		} else if soal == maxSoal {

			if skor < minSkor {
				minSkor = skor
				namaPemenang = nama
			}
		}
	}

	fmt.Println(namaPemenang, maxSoal, minSkor)
}
