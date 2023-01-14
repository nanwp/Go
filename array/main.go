package main

import "fmt"

func main(){
	scores := [8]int{100,80,75,92,70,93,88,67}
	fmt.Println("List nilai :",scores)

	/* Cari rata rata dari array di atas */
	var totalScore float64
	var avgScore float64 
	for _, score := range scores{
		totalScore = totalScore + float64(score)
	}
	avgScore = totalScore / float64(len(scores))
	fmt.Println("Total nilai :", totalScore, "\nRata - Rata :" ,avgScore)

	/*masukan nilai array di atas ke slice goodScore yang memiliki nilai >= 90*/

	var goodScore []int
	for _,score := range scores{
		if score >= 90{
			goodScore = append(goodScore, score)
		}
	}
	fmt.Println("Nilai Yang di atas 90",goodScore)


}

