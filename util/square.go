package util

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lijingbo8119/minesweeper-ebiten/core"
	"github.com/lijingbo8119/minesweeper-ebiten/resource"
)

func GetSquareImage(s *core.Square) *ebiten.Image {
	switch s.SquareStatus {
	case core.SquareStatusClosed:
		return resource.Resources.Blank
	case core.SquareStatusMouseDown:
		return resource.Resources.Mousedown
	case core.SquareStatusMarkedFlag:
		return resource.Resources.Flag
	case core.SquareStatusMarkedWrong:
		return resource.Resources.Mine1
	case core.SquareStatusOpened:
		switch len(s.AroundSquares.Filter(func(s *core.Square) bool { return s.SquareType == core.SquareTypeMine })) {
		case 0:
			return resource.Resources.MineNum0
		case 1:
			return resource.Resources.MineNum1
		case 2:
			return resource.Resources.MineNum2
		case 3:
			return resource.Resources.MineNum3
		case 4:
			return resource.Resources.MineNum4
		case 5:
			return resource.Resources.MineNum5
		case 6:
			return resource.Resources.MineNum6
		case 7:
			return resource.Resources.MineNum7
		case 8:
			return resource.Resources.MineNum8
		}
	case core.SquareStatusExploded:
		return resource.Resources.Mine2
	}
	return resource.Resources.Face0
}
