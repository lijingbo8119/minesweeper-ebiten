package core

import (
	"math/rand"
	"time"
)

type Matrix []Squares

func (this Matrix) FindSquare(closure func(s *Square) bool) *Square {
	for _, squares := range this {
		if find := squares.Find(closure); find != nil {
			return find
		}
	}
	return nil
}

func (this Matrix) FindSquares(closure func(s *Square) bool) Squares {
	res := Squares{}
	for _, row := range this {
		for _, s := range row {
			if closure(s) {
				res = append(res, s)
			}
		}
	}
	return res
}

func (this Matrix) MinesCount(squares Squares) int {
	count := 0
	for _, s := range squares {
		if s.SquareType == SquareTypeMine {
			count++
		}
	}
	return count
}

func NewMatrix(rowsLength int, colsLength int, minesCount int) *Matrix {
	if rowsLength < 9 || colsLength < 9 || minesCount < 10 {
		panic("初始化错误")
	}

	matrix := new(Matrix)

	tempArr := Squares{}
	squaresCount := rowsLength * colsLength
	for squaresCount > 0 {
		squareType := func() SquareType {
			if minesCount > 0 {
				minesCount--
				return SquareTypeMine
			}
			return SquareTypeNormal
		}()
		square := newSquare(squareType)
		tempArr = append(tempArr, square)
		squaresCount--
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(tempArr), func(i, j int) { tempArr[i], tempArr[j] = tempArr[j], tempArr[i] })

	index := 0
	for i := 0; i < rowsLength; i++ {
		*matrix = append(*matrix, make(Squares, colsLength, colsLength))
		for j := 0; j < colsLength; j++ {
			(*matrix)[i][j] = tempArr[index]
			(*matrix)[i][j].SquareCoordinate = NewCoordinate(i, j)
			index++
		}
	}

	for i := 0; i < rowsLength; i++ {
		for j := 0; j < colsLength; j++ {
			square := (*matrix)[i][j]
			square.AroundSquares = matrix.FindSquares(func(s *Square) bool { return square.SquareCoordinate.Near(s.SquareCoordinate) })
		}
	}

	return matrix
}
