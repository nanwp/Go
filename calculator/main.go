package main

import (
	"fmt"
	"calculator/calc"
)

func main(){

	
	// implemantasi tambah dan kurang
	nilai1 := calc.Tambah(10, 20)
	nilai2 := calc.Kurang(20, 5)
	nilai3 := calc.Kurang(nilai1, nilai2)
	fmt.Printf("Nilai 1 = %d Nilai 2 = %d Nilai 3 = %d \n", nilai1, nilai2, nilai3)

	//implementasi bagi //tipe data float
	nilai4 := calc.Bagi(20, 80)
	fmt.Printf("Hasil bagi adalah %f \n", nilai4)

	//implementasi kali
	fmt.Printf("Hasil perkalian adalah %d ", calc.Kali(2,9))

	
}