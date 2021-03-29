package game

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	update(game *Game)
	draw(screen *ebiten.Image, game *Game)
}
