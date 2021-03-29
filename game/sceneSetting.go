package game

import "github.com/hajimehoshi/ebiten/v2"

type Setting struct {
}

func (receiver Setting) update(game *Game) {}

func (receiver Setting) draw(screen *ebiten.Image, game *Game) {}
