package main

import (
	"errors"
	"image/color"
	"math"
	"math/rand"
)

type SalesPerson struct {
	entity

	movingTowards entity

	dx, dy float64
}

var NotPersonFound = errors.New("no person was found")

type VisitedSale struct {
	entity
}

func (*VisitedSale) Color() color.Color {
	return color.RGBA{0, 255, 0, 1}
}

var _ entity = &VisitedSale{}

func (p *SalesPerson) Update(w *World) error {
	if p.movingTowards != nil {
		p.SetXY(p.X() + p.dx, p.Y() + p.dy)

		if p.distanceFrom(p.movingTowards) <= 2 {
			w.Entities().Replace(p.movingTowards, &VisitedSale{p.movingTowards})
			p.movingTowards = nil
			p.dx = 0
			p.dy = 0
		}
		return nil
	}

	totalTimeInSeconds := rand.Float64() + 0.1

	nearestPerson := w.Entities().FindNearest(p, &Person{}, 1)
	if nearestPerson != nil {
		v := nearestPerson.(*Person)

		distance := p.distanceFrom(v)
		if distance <= rand.Float64() *100 {
			p.movingTowards = v
			p.dx = rateOfChange(p.X(), v.X(), totalTimeInSeconds)
			p.dy = rateOfChange(p.Y(), v.Y(), totalTimeInSeconds)
		}
	}

	return nil
}

func rateOfChange(s1 float64, s2 float64, t float64) float64 {
	velocity := (s2 - s1) / t
	return velocity / 60
}

func (*SalesPerson) Color() color.Color {
	return color.RGBA{255, 0, 0, 1}
}

func (p *SalesPerson) distanceFrom(e entity) float64 {
	dx := p.X() - e.X()
	dy := p.Y() - e.Y()
	return math.Sqrt(dx*dx + dy*dy)
}

var _ entity = &SalesPerson{}

func NewSalesPerson(w *World) entity {
	return &SalesPerson{entity: NewPerson(w)}
}
