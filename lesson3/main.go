package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type FaceitPlayer struct {
	Nickname    string
	Level       int
	ELO         int
	KDA         float64
	Kills       int
	Deaths      int
	Assists     int
	GamesPlayed int `json:"games_played"`
	Won         int
	Lost        int
	FPLGames    int
}

func (f *FaceitPlayer) Play() {
	// fmt.Printf("%+v", f)
	fmt.Printf("%s играет в CS2 сейчас", f.Nickname)
}

func (f *FaceitPlayer) AddKill() {
	f.Kills++
}

func (f *FaceitPlayer) AddDeath() {
	f.Deaths--
}

func (f *FaceitPlayer) WinGame(elo int) {
	f.GamesPlayed++
	f.Won++
	f.ELO += elo
}

func GetFaceitPlayerFromFile(filename string) (FaceitPlayer, error) {
	fd, err := os.ReadFile(filename)
	if err != nil {
		return FaceitPlayer{}, err
	}

	var faceit FaceitPlayer

	err = json.Unmarshal(fd, &faceit)
	if err != nil {
		return FaceitPlayer{}, err
	}

	return faceit, nil
}

func main() {
	faceit, err := GetFaceitPlayerFromFile("monesy.json")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(faceit)
	}
}
