package main

import "fmt"

func main() {
	//fmt.Println("test")

	/*
	buat string golang the best language, tampilkan jika index dari furufnya genap.
	*/
	kalimat := "golang the bset language"

	for index, letter := range kalimat{
		if index % 2 == 0 {
			fmt.Println(index,": Letter :",string(letter))
		}
	}

	/*
	cetak hanya huruf fokal
	*/

	for _, letter := range kalimat{
		switch string(letter){
		case "a","i","u","e","o":
			println("Vokal",string(letter))
		}
	}
}