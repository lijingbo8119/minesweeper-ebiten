package util

import (
	"github.com/gogf/gf/util/gconv"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lijingbo8119/minesweeper-ebiten/resource"
	"strings"
	"time"
)

func getNumberImage(n int) *ebiten.Image {
	if n >= 10 || n < 0 {
		panic("getNumberImage error: " + gconv.String(n))
	}
	switch n {
	case 0:
		return resource.Resources.Num0
	case 1:
		return resource.Resources.Num1
	case 2:
		return resource.Resources.Num2
	case 3:
		return resource.Resources.Num3
	case 4:
		return resource.Resources.Num4
	case 5:
		return resource.Resources.Num5
	case 6:
		return resource.Resources.Num6
	case 7:
		return resource.Resources.Num7
	case 8:
		return resource.Resources.Num8
	case 9:
		return resource.Resources.Num9
	}
	return resource.Resources.Num0
}

func DrawNumber(screen *ebiten.Image, offsetX, offsetY float64, n int) {
	nums := strings.Split(gconv.String(n), "")
	switch len(nums) {
	case 1:
		nums = []string{"0", "0", nums[0]}
	case 2:
		nums = []string{"0", nums[0], nums[1]}
	}

	op := &ebiten.DrawImageOptions{}

	for i, v := range nums {
		op.GeoM.Reset()
		op.GeoM.Translate(offsetX+gconv.Float64(i)*13, offsetY)
		screen.DrawImage(getNumberImage(gconv.Int(v)), op)
	}

}

func DrawNumberTime(screen *ebiten.Image, offsetX, offsetY float64, start *time.Time, end *time.Time) {
	var num int
	if start == nil {
		DrawNumber(screen, offsetX, offsetY, num)
		return
	}
	if end == nil {
		num = gconv.Int(time.Now().Sub(*start).Seconds())
	} else {
		num = gconv.Int(end.Sub(*start).Seconds())
	}
	DrawNumber(screen, offsetX, offsetY, num)
}
