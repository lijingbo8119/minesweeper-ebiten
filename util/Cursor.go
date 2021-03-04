package util

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lijingbo8119/minesweeper-ebiten/cursor"
	"github.com/lijingbo8119/minesweeper-ebiten/resource"
)

func GetCursorImage(action cursor.Action) *ebiten.Image {
	switch action {
	case cursor.ActionRelease:
		return resource.Resources.Cursor1
	case cursor.ActionPress:
		return resource.Resources.Cursor2
	}
	return resource.Resources.Cursor1
}
