package core

type SquareStatus int

const (
	SquareStatusClosed SquareStatus = iota
	SquareStatusMouseDown
	SquareStatusMarkedFlag
	SquareStatusMarkedWrong
	SquareStatusOpened
	SquareStatusExploded
)

type SquareType int

const (
	SquareTypeNormal SquareType = iota
	SquareTypeMine
)

type FaceStatus int

const (
	FaceStatusSmile FaceStatus = iota
	FaceStatusSmileMouseDown
	FaceStatusOps
	FaceStatusDied
	FaceStatusWin
)
