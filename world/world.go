package world

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kyroy/kdtree"
	"image"
)

type World struct {
	width, height int
	entities      entities
}

func (w *World) Update(screen *ebiten.Image) error {
	image := image.NewRGBA(image.Rect(0, 0, w.width, w.height))

	err := w.entities.Each(func(e Entity) error {
		err := e.Update(w)
		if err != nil {
			return err
		}

		image.Set(int(e.X()), int(e.Y()), e.Color())
		return nil
	})
	if err != nil {
		return err
	}

	screen.ReplacePixels(image.Pix)
	return nil
}

func (w *World) Entities() entities {
	return w.entities
}

func (w *World) Height() int {
	return w.height
}

func (w *World) Width() int {
	return w.width
}

func NewWorld(
	width, height int,
	inits []EntityInit,
) *World {
	world := &World{
		width:    width,
		height:   height,
		entities: map[string]*kdtree.KDTree{},
	}

	for _, entityInit := range inits {
		for i := 0; i < entityInit.Num; i++ {
			world.Entities().Add(entityInit.Init(world))
		}
	}

	return world
}
