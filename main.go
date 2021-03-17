package main

import (
	"embed"
	"github.com/gogf/gf/util/gconv"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lijingbo8119/minesweeper-ebiten/core"
	"github.com/lijingbo8119/minesweeper-ebiten/cursor"
	"github.com/lijingbo8119/minesweeper-ebiten/resource"
	"github.com/lijingbo8119/minesweeper-ebiten/util"
	_ "image/png"
	"log"
)

//go:embed images/*
var images embed.FS

const (
	MainBoardOffsetX = 10 / 2
	MainBoardOffsetY = (10 / 2) + 40
	screenWidth      = 480 + 10
	screenHeight     = 256 + 10 + 40
	faceOffsetX      = 232
	faceOffsetY      = 10
)

type Game struct {
	Matrix *core.Matrix
}

// 1/60 s
func (this *Game) Update() error {
	cursor.BindUpdate()
	return nil
}

// 1/60 s
func (this *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	util.DrawNumberTime(screen, 10, 10, core.State.GetStartTime(), core.State.GetEndTime())
	util.DrawNumber(screen, screenWidth-50, 10, core.State.UnmarkedMinesCount)

	op.GeoM.Reset()
	op.GeoM.Translate(faceOffsetX, faceOffsetY)
	screen.DrawImage(util.GetFaceImage(core.State.Face), op)

	core.State.Matrix.
		FindSquares(func(s *core.Square) bool { return true }).
		Each(func(s *core.Square) {
			op.GeoM.Reset()
			x, y := util.Coordinate2Position(s.SquareCoordinate.I, s.SquareCoordinate.J)
			op.GeoM.Translate(MainBoardOffsetX, MainBoardOffsetY)
			op.GeoM.Translate(x, y)
			screen.DrawImage(util.GetSquareImage(s), op)
		})

	op.GeoM.Reset()
	cursorPosition := cursor.GetPosition()
	x, y := cursorPosition.X, cursorPosition.Y
	op.GeoM.Translate(-2, -2)
	op.GeoM.Translate(gconv.Float64(x), gconv.Float64(y))
	//screen.DrawImage(util.GetCursorImage(core.State.CursorAction), op)

	ebitenutil.DebugPrint(screen, gconv.String(gconv.Int(ebiten.CurrentTPS())))
}

func (this *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

var game = &Game{}

func main() {
	resource.Init(images)

	//ebiten.SetCursorMode(ebiten.CursorModeHidden)

	core.State.Start(30, 16, 99)

	cursor.RegisterEvent(ebiten.MouseButtonLeft, cursor.ActionPress, func(s *cursor.Status) bool {
		core.State.CursorAction = cursor.ActionPress
		if util.IsCursorInWidget(faceOffsetX, faceOffsetY, 26, 26, *s) {
			core.State.Face.SetStatus(core.FaceStatusSmileMouseDown)
			return true
		}
		return true
	})

	cursor.RegisterEvent(ebiten.MouseButtonLeft, cursor.ActionRelease, func(s *cursor.Status) bool {
		core.State.CursorAction = cursor.ActionRelease
		if util.IsCursorInWidget(faceOffsetX, faceOffsetY, 26, 26, *s) && core.State.Face.FaceStatus == core.FaceStatusSmileMouseDown {
			core.State.Start(30, 16, 99)
			return true
		}
		if !util.IsCursorInWidget(faceOffsetX, faceOffsetY, 26, 26, *s) && core.State.Face.FaceStatus == core.FaceStatusSmileMouseDown {
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
		i, j := util.Position2Coordinate(gconv.Float64(s.Position.X-MainBoardOffsetX), gconv.Float64(s.Position.Y-MainBoardOffsetY))
		if i == -1 || j == -1 {
			return false
		}
		core.State.MouseState.LeftMouseDown(core.NewCoordinate(i, j))
		return true
	})

	cursor.RegisterEvent(ebiten.MouseButtonLeft, cursor.ActionRelease, func(s *cursor.Status) bool {
		core.State.CursorAction = cursor.ActionRelease
		i, j := util.Position2Coordinate(gconv.Float64(s.Position.X-MainBoardOffsetX), gconv.Float64(s.Position.Y-MainBoardOffsetY))
		if i == -1 || j == -1 {
			return false
		}
		core.State.MouseState.LeftMouseUp(core.NewCoordinate(i, j))
		return true
	})

	cursor.RegisterEvent(ebiten.MouseButtonRight, cursor.ActionPress, func(s *cursor.Status) bool {
		core.State.CursorAction = cursor.ActionPress
		i, j := util.Position2Coordinate(gconv.Float64(s.Position.X-MainBoardOffsetX), gconv.Float64(s.Position.Y-MainBoardOffsetY))
		if i == -1 || j == -1 {
			return false
		}
		core.State.MouseState.RightMouseDown(core.NewCoordinate(i, j))
		return true
	})

	cursor.RegisterEvent(ebiten.MouseButtonRight, cursor.ActionRelease, func(s *cursor.Status) bool {
		core.State.CursorAction = cursor.ActionRelease
		i, j := util.Position2Coordinate(gconv.Float64(s.Position.X-MainBoardOffsetX), gconv.Float64(s.Position.Y-MainBoardOffsetY))
		if i == -1 || j == -1 {
			return false
		}
		core.State.MouseState.RightMouseUp(core.NewCoordinate(i, j))
		return true
	})

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Minesweeper")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
