package main

import "fmt"

type Gamer struct {
	Name string
	Games []string
}

func (gamer *Gamer) AddGame(game string){
	gamer.Games = append(gamer.Games, game)
}

func main(){
	nanda := Gamer{Name: "Nanda"}
	nanda.AddGame("GTA")
	nanda.AddGame("ML")
	nanda.AddGame("PUBG")

	for _, games := range nanda.Games{

		fmt.Println(games)

	}

}