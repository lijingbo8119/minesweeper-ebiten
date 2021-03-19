package core

type Square struct {
	SquareStatus     SquareStatus
	SquareType       SquareType
	SquareCoordinate Coordinate
	AroundSquares    Squares
}

func (this *Square) setStatus(s SquareStatus) {
	this.SquareStatus = s
}

func (this *Square) open(triggeredByClick bool) bool {
	if this.SquareStatus != SquareStatusClosed && this.SquareStatus != SquareStatusMouseDown {
		return true
	}
	if this.SquareType == SquareTypeMine && triggeredByClick {
		this.setStatus(SquareStatusExploded)
		return false
	}

	unmarkedMines := this.AroundSquares.Filter(func(s *Square) bool {
		return s.SquareType == SquareTypeMine && s.SquareStatus != SquareStatusMarkedFlag
	})
	if triggeredByClick {
		if len(unmarkedMines) == 0 {
			this.openAroundSquares()
		}
		this.setStatus(SquareStatusOpened)
	}

	if !triggeredByClick && this.SquareType == SquareTypeNormal {
		this.setStatus(SquareStatusOpened)
	}

	if !triggeredByClick && this.SquareType == SquareTypeNormal && len(unmarkedMines) == 0 {
		this.openAroundSquares()
	}

	return true
}

func (this *Square) mark() {
	if this.SquareStatus != SquareStatusClosed && this.SquareStatus != SquareStatusMouseDown && this.SquareStatus != SquareStatusMarkedFlag {
		return
	}
	if this.SquareStatus == SquareStatusMarkedFlag {
		this.setStatus(SquareStatusClosed)
		return
	}
	this.setStatus(SquareStatusMarkedFlag)
}

func (this *Square) openAroundSquares() bool {
	markedWrongMines := this.AroundSquares.
		Filter(func(s *Square) bool { return s.SquareType == SquareTypeNormal && s.SquareStatus == SquareStatusMarkedFlag })
	if len(markedWrongMines) > 0 {
		markedWrongMines.Each(func(s *Square) { s.setStatus(SquareStatusMarkedWrong) })
		this.AroundSquares.
			Filter(func(s *Square) bool { return s.SquareType == SquareTypeMine && (s.SquareStatus == SquareStatusClosed || s.SquareStatus == SquareStatusMouseDown) }).
			Each(func(s *Square) { s.setStatus(SquareStatusOpened) })
		return false
	}

	if unmarkedMines := this.AroundSquares.Filter(func(s *Square) bool {
		return s.SquareType == SquareTypeMine && s.SquareStatus != SquareStatusMarkedFlag
	}); len(unmarkedMines) > 0 {
		return true
	}
	for _, s := range this.AroundSquares {
		s.open(false)
	}
	return true
}

func newSquare(t SquareType) *Square {
	s := &Square{
		SquareType: t,
	}
	s.setStatus(SquareStatusClosed)
	return s
}
