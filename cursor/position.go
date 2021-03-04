package cursor

import "github.com/hajimehoshi/ebiten/v2"

type Position struct {
	X, Y int
}

func GetPosition() Position {
	x, y := ebiten.CursorPosition()
	return Position{x, y}
}
