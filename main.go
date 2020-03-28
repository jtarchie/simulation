package main

import (
	"github.com/hajimehoshi/ebiten"
	"log"
)

func main() {
	world := NewWorld(
		500,
		500,
		[]entityInit{
			{1000, NewPerson},
			{100, NewSalesPerson},
		},
	)
	err := ebiten.Run(
		world.Update,
		world.width,
		world.height,
		2,
		"Simulation",
	)
	if err != nil {
		log.Fatal(err)
	}
}
