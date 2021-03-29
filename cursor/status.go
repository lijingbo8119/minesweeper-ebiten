package cursor

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Status struct {
	Button   ebiten.MouseButton
	Position Position
	Action   Action
}

func newStatus(b ebiten.MouseButton, p Position, a Action) *Status {
	return &Status{Button: b, Position: p, Action: a}
}
