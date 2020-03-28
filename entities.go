package main

import (
	"fmt"
	"github.com/kyroy/kdtree"
	"image/color"
)

type entityInit struct {
	num int
	init func(*World) entity
}

type entity interface {
	kdtree.Point

	X() float64
	Y() float64
	SetXY(float64, float64)
	Color() color.Color
	Update(*World) error
}

type entities map[string]*kdtree.KDTree

func (e entities) get(original entity) *kdtree.KDTree {
	t := fmt.Sprintf("%T", original)
	if _, ok := e[t]; !ok {
		e[t] = kdtree.New([]kdtree.Point{})
	}
	return e[t]
}

func (e entities) Replace(original, with entity) {
	e.get(original).Remove(original)
	e.Add(with)
}

func (e entities) Add(original entity) {
	t := fmt.Sprintf("%T", original)
	if tree, ok := e[t]; ok {
		tree.Insert(original)
	} else {
		e[t] = kdtree.New([]kdtree.Point{original})
	}
}

func (e entities) Each(f func(entity) error) error {
	for _, tree := range e {
		for _, point := range tree.Points() {
			entity := point.(entity)
			err := f(entity)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (e entities) FindNearest(from entity, of entity, i int) interface{} {
	points := e.get(of).KNN(from, 1)

	if len(points) == 0 {
		return nil
	}

	return points[0]
}
