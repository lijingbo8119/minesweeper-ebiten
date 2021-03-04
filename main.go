package main

import (
	"ebientest/core"
	"ebientest/cursor"
	"ebientest/resource"
	"ebientest/util"
	"embed"
	"github.com/gogf/gf/util/gconv"
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	"log"
)

//go:embed images/*
var images embed.FS

const (
	MainBoardOffsetX = 4
	MainBoardOffsetY = 9
	screenWidth      = 320
	screenHeight     = 240
)

type Game struct {
	Matrix *core.Matrix
}

func (this *Game) Update() error {
	cursor.BindUpdate()
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
}

func (this *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

var game = &Game{}

func main() {
	resource.Init(images)

	core.State.SetMatrixParam(10, 10, 10)

	cursor.RegisterEvent(ebiten.MouseButtonLeft, cursor.ActionPress, func(s *cursor.Status) {
		i, j := util.Position2Coordinate(gconv.Float64(s.Position.X-MainBoardOffsetX), gconv.Float64(s.Position.Y-MainBoardOffsetY))
		core.State.MouseState.LeftMouseDown(core.NewCoordinate(i, j))
	})

	cursor.RegisterEvent(ebiten.MouseButtonLeft, cursor.ActionRelease, func(s *cursor.Status) {
		i, j := util.Position2Coordinate(gconv.Float64(s.Position.X-MainBoardOffsetX), gconv.Float64(s.Position.Y-MainBoardOffsetY))
		core.State.MouseState.LeftMouseUp(core.NewCoordinate(i, j))
	})

	cursor.RegisterEvent(ebiten.MouseButtonRight, cursor.ActionPress, func(s *cursor.Status) {
		i, j := util.Position2Coordinate(gconv.Float64(s.Position.X-MainBoardOffsetX), gconv.Float64(s.Position.Y-MainBoardOffsetY))
		core.State.MouseState.RightMouseDown(core.NewCoordinate(i, j))
	})

	cursor.RegisterEvent(ebiten.MouseButtonRight, cursor.ActionRelease, func(s *cursor.Status) {
		i, j := util.Position2Coordinate(gconv.Float64(s.Position.X-MainBoardOffsetX), gconv.Float64(s.Position.Y-MainBoardOffsetY))
		core.State.MouseState.RightMouseUp(core.NewCoordinate(i, j))
	})

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Keyboard (Ebiten Demo)")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
