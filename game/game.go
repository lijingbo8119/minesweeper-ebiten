package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lijingbo8119/minesweeper-ebiten/cursor"
)

const (
	ScreenWidth  = 480 + 10
	ScreenHeight = 256 + 10 + 40
)

type Game struct {
	Scene Scene
}

// 1/60 s
func (this *Game) Update() error {
	cursor.CleanAllEvent()
	this.Scene.update(this)
	cursor.BindUpdate()
	return nil
}

// 1/60 s
func (this *Game) Draw(screen *ebiten.Image) {
	this.Scene.draw(screen, this)
}

func (this *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
