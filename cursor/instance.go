package cursor

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var leftButtonValInstance *Val

var rightButtonValInstance *Val

func BindUpdate() {
	if leftButtonValInstance == nil {
		leftButtonValInstance = NewVal(newStatus(ebiten.MouseButtonLeft, GetPosition(), ActionRelease))
	}

	if rightButtonValInstance == nil {
		rightButtonValInstance = NewVal(newStatus(ebiten.MouseButtonRight, GetPosition(), ActionRelease))
	}

	GetAction(ebiten.MouseButtonLeft)
	GetAction(ebiten.MouseButtonRight)
}
