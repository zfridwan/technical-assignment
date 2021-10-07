package main

import "fmt"

func main() {
	soal1(10)
	fmt.Println()
}

func soal1(angka int) {

	//definisi variable dan aritmatikas matematika pertambahan
	var var_angka = (angka * 2) + 1

	for i := 1; i <= var_angka; i++ {
		
		for j := 1; j <= var_angka; j++ {
			if (i == 1 || j == 1) || (i == var_angka || j == var_angka) || (i == angka+1 || j == angka+1) || (j == (var_angka-i+1) || i == j) {
				fmt.Print("x ")
			} else {
				fmt.Print("  ")
			}

		}
		fmt.Println()
	}

}