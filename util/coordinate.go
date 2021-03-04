package util

import (
	"github.com/gogf/gf/util/gconv"
	"math"
)

const squareSize = 16

func Coordinate2Position(i, j int) (x, y float64) {
	x = gconv.Float64(i * squareSize)
	y = gconv.Float64(j * squareSize)
	return
}

func Position2Coordinate(x, y float64) (i, j int) {
	if x < 0 || y < 0 {
		return -1, -1
	}
	return gconv.Int(math.Floor(x / squareSize)), gconv.Int(math.Floor(y / squareSize))
}
