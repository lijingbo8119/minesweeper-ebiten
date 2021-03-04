package core_test

import (
	"github.com/lijingbo8119/minesweeper-ebiten/core"
	"fmt"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	matrix := core.NewMatrix(10, 10, 10)
	fmt.Println(matrix)
}
