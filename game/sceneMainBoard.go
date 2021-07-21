package game

import (
	"github.com/gogf/gf/util/gconv"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lijingbo8119/minesweeper-ebiten/core"
	"github.com/lijingbo8119/minesweeper-ebiten/cursor"
	"github.com/lijingbo8119/minesweeper-ebiten/util"
)

const (
	mainBoardOffsetX = 10 / 2
	mainBoardOffsetY = (10 / 2) + 40

	mainBoardFaceIconOffsetX = 232
	mainBoardFaceIconOffsetY = 10

	mainBoardFaceIconWidth  = 26
	mainBoardFaceIconHeight = 26
)

type MainBoard struct{}

func (receiver MainBoard) update(game *Game) {
	cursor.RegisterEvent(ebiten.MouseButtonLeft, cursor.ActionPress, func(s *cursor.Status) bool {
		core.State.CursorAction = cursor.ActionPress
		if util.IsCursorInWidget(mainBoardFaceIconOffsetX, mainBoardFaceIconOffsetY, mainBoardFaceIconWidth, mainBoardFaceIconHeight, *s) {
			core.State.Face.SetStatus(core.FaceStatusSmileMouseDown)
			return true
		}
		return true
	})

	cursor.RegisterEvent(ebiten.MouseButtonLeft, cursor.ActionRelease, func(s *cursor.Status) bool {
		core.State.CursorAction = cursor.ActionRelease
		if util.IsCursorInWidget(mainBoardFaceIconOffsetX, mainBoardFaceIconOffsetY, mainBoardFaceIconWidth, mainBoardFaceIconHeight, *s) &&
			core.State.Face.FaceStatus == core.FaceStatusSmileMouseDown {
			core.State.Start(30, 16, 99)
			return true
		}
		if !util.IsCursorInWidget(mainBoardFaceIconOffsetX, mainBoardFaceIconOffsetY, mainBoardFaceIconWidth, mainBoardFaceIconHeight, *s) &&
			core.State.Face.FaceStatus == core.FaceStatusSmileMouseDown {
			if core.State.GetEndTime() != nil {
				core.State.Face.SetStatus(core.FaceStatusDied)
			} else {
				core.State.Face.SetStatus(core.FaceStatusSmile)
			}
		}
		return true
	})

	cursor.RegisterEvent(ebiten.MouseButtonRight, cursor.ActionRelease, func(s *cursor.Status) bool {
		core.State.CursorAction = cursor.ActionRelease
		if util.IsCursorInWidget(mainBoardFaceIconOffsetX, mainBoardFaceIconOffsetY, mainBoardFaceIconWidth, mainBoardFaceIconHeight, *s) {
			game.Scene = Setting{}
			return true
		}
		if !util.IsCursorInWidget(mainBoardFaceIconOffsetX, mainBoardFaceIconOffsetY, mainBoardFaceIconWidth, mainBoardFaceIconHeight, *s) &&
			core.State.Face.FaceStatus == core.FaceStatusSmileMouseDown {
			if core.State.GetEndTime() != nil {
				core.State.Face.SetStatus(core.FaceStatusDied)
			} else {
				core.State.Face.SetStatus(core.FaceStatusSmile)
			}
		}
		return true
	})

	cursor.RegisterEvent(ebiten.MouseButtonLeft, cursor.ActionPress, func(s *cursor.Status) bool {
		core.State.CursorAction = cursor.ActionPress
		i, j := util.Position2Coordinate(gconv.Float64(s.Position.X-mainBoardOffsetX), gconv.Float64(s.Position.Y-mainBoardOffsetY))
		if i == -1 || j == -1 {
			return false
		}
		core.State.MouseState.LeftMouseDown(core.NewCoordinate(i, j))
		return true
	})

	cursor.RegisterEvent(ebiten.MouseButtonLeft, cursor.ActionRelease, func(s *cursor.Status) bool {
		core.State.CursorAction = cursor.ActionRelease
		i, j := util.Position2Coordinate(gconv.Float64(s.Position.X-mainBoardOffsetX), gconv.Float64(s.Position.Y-mainBoardOffsetY))
		if i == -1 || j == -1 {
			return false
		}
		core.State.MouseState.LeftMouseUp(core.NewCoordinate(i, j))
		return true
	})

	cursor.RegisterEvent(ebiten.MouseButtonRight, cursor.ActionPress, func(s *cursor.Status) bool {
		core.State.CursorAction = cursor.ActionPress
		i, j := util.Position2Coordinate(gconv.Float64(s.Position.X-mainBoardOffsetX), gconv.Float64(s.Position.Y-mainBoardOffsetY))
		if i == -1 || j == -1 {
			return false
		}
		core.State.MouseState.RightMouseDown(core.NewCoordinate(i, j))
		return true
	})

	cursor.RegisterEvent(ebiten.MouseButtonRight, cursor.ActionRelease, func(s *cursor.Status) bool {
		core.State.CursorAction = cursor.ActionRelease
		i, j := util.Position2Coordinate(gconv.Float64(s.Position.X-mainBoardOffsetX), gconv.Float64(s.Position.Y-mainBoardOffsetY))
		if i == -1 || j == -1 {
			return false
		}
		core.State.MouseState.RightMouseUp(core.NewCoordinate(i, j))
		return true
	})
}

func (receiver MainBoard) draw(screen *ebiten.Image, game *Game) {
	op := &ebiten.DrawImageOptions{}

	util.DrawNumberTime(screen, 10, 10, core.State.GetStartTime(), core.State.GetEndTime())
	util.DrawNumber(screen, ScreenWidth-50, 10, core.State.UnmarkedMinesCount)

	op.GeoM.Reset()
	op.GeoM.Translate(mainBoardFaceIconOffsetX, mainBoardFaceIconOffsetY)
	screen.DrawImage(util.GetFaceImage(core.State.Face), op)

	core.State.Matrix.
		FindSquares(func(s *core.Square) bool { return true }).
		Each(func(s *core.Square) {
			op.GeoM.Reset()
			x, y := util.Coordinate2Position(s.SquareCoordinate.I, s.SquareCoordinate.J)
			op.GeoM.Translate(mainBoardOffsetX, mainBoardOffsetY)
			op.GeoM.Translate(x, y)
			screen.DrawImage(util.GetSquareImage(s), op)
		})

	op.GeoM.Reset()
	cursorPosition := cursor.GetPosition()
	x, y := cursorPosition.X, cursorPosition.Y
	op.GeoM.Translate(-2, -2)
	op.GeoM.Translate(gconv.Float64(x), gconv.Float64(y))

	ebitenutil.DebugPrint(screen, gconv.String(gconv.Int(ebiten.CurrentTPS())))
}
