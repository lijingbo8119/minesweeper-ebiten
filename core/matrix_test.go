package core_test

import (
	"ebientest/core"
	"fmt"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	matrix := core.NewMatrix(10, 10, 10)
	fmt.Println(matrix)
}
