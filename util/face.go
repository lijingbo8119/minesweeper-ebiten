package util

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lijingbo8119/minesweeper-ebiten/core"
	"github.com/lijingbo8119/minesweeper-ebiten/resource"
)

func GetFaceImage(f *core.Face) *ebiten.Image {
	switch f.FaceStatus {
	case core.FaceStatusSmile:
		return resource.Resources.Face0
	case core.FaceStatusSmileMouseDown:
		return resource.Resources.Face1
	case core.FaceStatusOps:
		return resource.Resources.Face2
	case core.FaceStatusDied:
		return resource.Resources.Face3
	case core.FaceStatusWin:
		return resource.Resources.Face4
	default:
		return resource.Resources.Face0
	}
}


