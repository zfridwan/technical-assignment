package main

import "fmt"

func main() {
	//Definisi tipe data dan variable
	angka_A := 17
	angka_B := 8
	soal3(angka_A, angka_B)
}

func soal3(a int, b int) {
	//operator aritmatika pertambahan dari variable a tambah
	jumlah := a + b

	for i := 1; i <= jumlah; i++ {

		if i%3 == 0 {
			fmt.Printf("b")
		} else {
			fmt.Printf("a")
		}

	}
}