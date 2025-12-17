package main

import (
	"fmt"
)

func main() {
	var nilai int
	var valid bool
	var grade string
	var ulang string
for {
	
	fmt.Print("Masukkan Nilai :")
	fmt.Scan(&nilai)
	
	valid = nilai < 0 || nilai > 100 ;
	if valid {
		fmt.Println("Masukkan tidak valid! Ulangi lagi")
		continue
	}
	
	if nilai >= 90 {
		grade = "A"
		}else if nilai >= 80 {
			grade = "B"
			}else if nilai >= 70  {
				grade = "C"
				}else if nilai >= 60 {
					grade = "D"
					}else {
		grade ="E"
	}
	fmt.Println("GRADE :" + grade)

	fmt.Println("Apakah anda ingin mengulangi program? (y/n)")
	fmt.Scan(&ulang)
	if ulang != "Y" && ulang != "y" {
		break
	}
}
}