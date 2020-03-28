package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/jtarchie/simulation/people"
	 "github.com/jtarchie/simulation/world"
	"log"
)

func main() {
	w := world.NewWorld(
		500,
		500,
		[]world.EntityInit{
			{1000, people.NewPerson},
			{100, people.NewSalesPerson},
		},
	)
	err := ebiten.Run(
		w.Update,
		w.Width(),
		w.Height(),
		2,
		"Simulation",
	)
	if err != nil {
		log.Fatal(err)
	}
}
