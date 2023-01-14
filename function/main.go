package main

import (
	"errors"
	"fmt"
)

func main() {
	scores := []int{10, 5, 8, 9, 7}

	total := sum(scores)

	fmt.Println(total)
	
	result, eer := calculate(10, 2 , "=")

	if eer != nil{
		fmt.Println("Pesan : ",eer.Error())
	} else{
		fmt.Println("Hasil : ", result)
	}

}

func sum(scores []int) (total int) {

	for _, score := range scores {
		total = total + score
	}
	return 
}

func calculate(nilai1 int, nilai2 int, opsi string) (hasil int, errorResoult error){


	switch opsi {
	case "+":
		hasil = nilai1 + nilai2
	case "-":
		hasil = nilai1 - nilai2
	case "*":
		hasil = nilai1 * nilai2
	case "/":
		hasil = nilai1 / nilai2
	default:
		errorResoult = errors.New("Unknow opration")
	}

	return

}