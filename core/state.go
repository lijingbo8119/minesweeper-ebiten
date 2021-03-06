package core

import (
	"github.com/lijingbo8119/minesweeper-ebiten/cursor"
	"time"
)

type _state struct {
	CursorAction cursor.Action

	MouseState *MouseState

	UnmarkedMinesCount int
	Matrix             *Matrix

	Face *Face

	startTime *time.Time
	endTime   *time.Time
}

func (this *_state) GetStartTime() *time.Time {
	return this.startTime
}

func (this *_state) SetStartTime(t ...*time.Time) {
	if len(t) > 0 {
		this.startTime = t[0]
	} else {
		_t := time.Now()
		this.startTime = &_t
	}
}

func (this *_state) GetEndTime() *time.Time {
	return this.endTime
}

func (this *_state) SetEndTime(t ...*time.Time) {
	if len(t) > 0 {
		this.endTime = t[0]
	} else {
		_t := time.Now()
		this.endTime = &_t
	}
}

func (this *_state) Start(rowsLength int, colsLength int, minesCount int) {
	this.Face = &Face{FaceStatusSmile}

	this.CursorAction = cursor.ActionRelease

	this.UnmarkedMinesCount = minesCount
	this.Matrix = NewMatrix(rowsLength, colsLength, minesCount)
	this.SetStartTime(nil)
	this.SetEndTime(nil)
	this.MouseState = NewMouseState().
		RegisterLeftMouseDownHandler(this.leftMouseDownHandler).RegisterLeftMouseUpHandler(this.leftMouseUpHandler).
		RegisterRightMouseDownHandler(this.rightMouseDownHandler).RegisterRightMouseUpHandler(this.rightMouseUpHandler).
		RegisterLeftRightMouseDownHandler(this.leftRightMouseDownHandler).RegisterLeftRightMouseUpHandler(this.leftRightMouseUpHandler).
		RegisterResetHandler(this.resetHandler)
}

func (this *_state) resetHandler(c Coordinate) {
	this.Matrix.
		FindSquares(func(square *Square) bool { return square.SquareStatus == SquareStatusMouseDown }).
		Each(func(s *Square) { s.setStatus(SquareStatusClosed) })
}

func (this *_state) leftMouseDownHandler(c Coordinate) {
	if this.endTime != nil {
		return
	}
	this.Face.SetStatus(FaceStatusOps)
	this.Matrix.
		FindSquares(func(square *Square) bool { return square.SquareCoordinate.Equal(c) && square.SquareStatus == SquareStatusClosed }).
		Each(func(s *Square) { s.setStatus(SquareStatusMouseDown) })
}

func (this *_state) leftMouseUpHandler(c Coordinate) {
	if this.endTime != nil {
		return
	}
	if this.startTime == nil {
		this.SetStartTime()
	}
	this.Face.SetStatus(FaceStatusSmile)
	this.Matrix.
		FindSquares(func(square *Square) bool { return square.SquareCoordinate.Equal(c) && square.SquareStatus == SquareStatusMouseDown }).
		Each(func(s *Square) {
			if res := s.open(true); !res {
				this.SetEndTime()
				this.Face.SetStatus(FaceStatusDied)
			}
		})
}

func (this *_state) rightMouseDownHandler(c Coordinate) {
	if this.endTime != nil {
		return
	}
	this.Face.SetStatus(FaceStatusOps)
}

func (this *_state) rightMouseUpHandler(c Coordinate) {
	if this.endTime != nil {
		return
	}
	if this.startTime == nil {
		this.SetStartTime()
	}
	this.Face.SetStatus(FaceStatusSmile)
	this.Matrix.
		FindSquares(func(square *Square) bool { return square.SquareCoordinate.Equal(c) }).
		Each(func(s *Square) { s.mark() })
	this.UnmarkedMinesCount = len(this.Matrix.FindSquares(func(square *Square) bool { return square.SquareType == SquareTypeMine })) - len(this.Matrix.FindSquares(func(square *Square) bool { return square.SquareStatus == SquareStatusMarkedFlag }))
}

func (this *_state) leftRightMouseDownHandler(c Coordinate) {
	if this.endTime != nil {
		return
	}
	square := this.Matrix.FindSquare(func(square *Square) bool { return square.SquareCoordinate.Equal(c) })
	if square == nil {
		return
	}

	switch square.SquareStatus {
	case SquareStatusMouseDown:
		break
	case SquareStatusOpened:
		break
	case SquareStatusClosed:
		return
	default:
		return
	}

	square.AroundSquares.
		Filter(func(s *Square) bool {
			return s.SquareStatus == SquareStatusClosed
		}).
		Each(func(s *Square) {
			s.setStatus(SquareStatusMouseDown)
		})
}

func (this *_state) leftRightMouseUpHandler(c Coordinate) {
	if this.endTime != nil {
		return
	}
	square := this.Matrix.FindSquare(func(square *Square) bool { return square.SquareCoordinate.Equal(c) })
	if square == nil {
		return
	}

	switch square.SquareStatus {
	case SquareStatusOpened:
		if !square.openAroundSquares() {
			this.SetEndTime()
			this.Face.SetStatus(FaceStatusDied)
		}
		break
	default:
		return
	}
}

var State = new(_state)
