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
	"time"
)

//go:embed images/*
var images embed.FS

const (
	MainBoardOffsetX = 10 / 2
	MainBoardOffsetY = 10 / 2
	screenWidth      = 480 + 10
	screenHeight     = 256 + 10
)

type Game struct {
	Matrix *core.Matrix
}

func (this *Game) Update() error {
	cursor.BindUpdate()
	if endTime := core.State.GetEndTime(); endTime != nil && time.Now().Sub(*endTime) > 1*time.Second {
		core.State.SetMatrixParam(30, 16, 99)
	}
	return nil
}

func (this *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

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
	screen.DrawImage(util.GetCursorImage(core.State.CursorAction), op)

	ebitenutil.DebugPrint(screen, gconv.String(gconv.Int(ebiten.CurrentTPS())))
}

func (this *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

var game = &Game{}

func main() {
	resource.Init(images)

	ebiten.SetCursorMode(ebiten.CursorModeHidden)

	core.State.SetMatrixParam(30, 16, 99)

	cursor.RegisterEvent(ebiten.MouseButtonLeft, cursor.ActionPress, func(s *cursor.Status) {
		core.State.CursorAction = cursor.ActionPress
		i, j := util.Position2Coordinate(gconv.Float64(s.Position.X-MainBoardOffsetX), gconv.Float64(s.Position.Y-MainBoardOffsetY))
		core.State.MouseState.LeftMouseDown(core.NewCoordinate(i, j))
	})

	cursor.RegisterEvent(ebiten.MouseButtonLeft, cursor.ActionRelease, func(s *cursor.Status) {
		core.State.CursorAction = cursor.ActionRelease
		i, j := util.Position2Coordinate(gconv.Float64(s.Position.X-MainBoardOffsetX), gconv.Float64(s.Position.Y-MainBoardOffsetY))
		core.State.MouseState.LeftMouseUp(core.NewCoordinate(i, j))
	})

	cursor.RegisterEvent(ebiten.MouseButtonRight, cursor.ActionPress, func(s *cursor.Status) {
		core.State.CursorAction = cursor.ActionPress
		i, j := util.Position2Coordinate(gconv.Float64(s.Position.X-MainBoardOffsetX), gconv.Float64(s.Position.Y-MainBoardOffsetY))
		core.State.MouseState.RightMouseDown(core.NewCoordinate(i, j))
	})

	cursor.RegisterEvent(ebiten.MouseButtonRight, cursor.ActionRelease, func(s *cursor.Status) {
		core.State.CursorAction = cursor.ActionRelease
		i, j := util.Position2Coordinate(gconv.Float64(s.Position.X-MainBoardOffsetX), gconv.Float64(s.Position.Y-MainBoardOffsetY))
		core.State.MouseState.RightMouseUp(core.NewCoordinate(i, j))
	})

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Minesweeper")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
