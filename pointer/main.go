package main

import "fmt"

//membuat struct
type Gamer struct { 
	Name string
	Games []string
}

// menambah asterik pada tipe gamer sebagai keluaran agar tidak berubah
func (gamer *Gamer) AddGame(game string){ 
	gamer.Games = append(gamer.Games, game)
}


// main func
func main(){ 
	nanda := Gamer{Name: "Nanda"}
	nanda.AddGame("GTA")
	nanda.AddGame("ML")
	nanda.AddGame("PUBG")

	for _, games := range nanda.Games{

		fmt.Println(games)

	}

}