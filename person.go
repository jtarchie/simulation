package main

import (
	"image/color"
	"math/rand"
)

type Person struct {
	x, y float64
}

func (p *Person) SetXY(x float64, y float64) {
	p.x = x
	p.y = y
}

func (*Person) Dimensions() int {
	return 2
}

func (p *Person) Dimension(i int) float64 {
	if i == 1 {
		return p.x
	}
	return p.y
}

func (p *Person) Update(*World) error {
	return nil
}

func (p *Person) X() float64 {
	return p.x
}

func (p *Person) Y() float64 {
	return p.y
}

func (p *Person) Color() color.Color {
	return color.RGBA{255, 255, 255, 1}
}

func NewPerson(world *World) entity {
	return &Person{
		float64(rand.Intn(world.width)),
		float64(rand.Intn(world.height)),
	}
}

var _ entity = &Person{}
