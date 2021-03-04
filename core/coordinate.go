package core

import (
	"math"
)

type Coordinate struct {
	I int
	J int
}

func (this Coordinate) Equal(c Coordinate) bool {
	return !this.IsEmpty() && this.I == c.I && this.J == c.J
}

func (this Coordinate) Near(c Coordinate) bool {
	if this.Equal(c) {
		return false
	}
	return int(math.Abs(float64(this.I-c.I))) <= 1 && int(math.Abs(float64(this.J-c.J))) <= 1
}

func (this Coordinate) IsEmpty() bool {
	return this.I == -1 && this.J == -1
}

func NewCoordinate(i int, j int) Coordinate {
	return Coordinate{i, j}
}

func NewEmptyCoordinate() Coordinate {
	return Coordinate{-1, -1}
}
