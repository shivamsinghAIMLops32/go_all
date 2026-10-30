package main

import (
	"fmt"
	"time"
)

type Player struct {
	ID         int
	Name       string
	TimePlayed time.Time
}

func newPlayer(id int, name string) *Player {
	return &Player{
		ID:         id,
		Name:       name,
		TimePlayed: time.Now(),
	}
}

type Game struct {
	ID                int
	Name              string
	DeviceTypeSupport []string
	Players           []*Player
	PlayerScore       map[string]int // we need to make it cause without making if we try  to write anything it will panic
}

func NewGame(name string, devices []string, playersName []string) *Game {
	plList := make([]*Player, 0, len(playersName))
	mp := make(map[string]int, len(playersName))
	for i, name := range playersName {
		plList = append(plList, newPlayer(i, name))
		mp[name] = i * 10
	}

	return &Game{
		ID:                1,
		Name:              name,
		DeviceTypeSupport: devices,
		Players:           plList,
		PlayerScore:       mp,
	}
}

func main() {
	// struct get assined with 0 values if not given
	// also when we dont write keys then we have to give values and that too in order
	game := NewGame("pubg", []string{"ios", "android"}, []string{"shroud", "drdoom"})
	fmt.Println(game)
}
