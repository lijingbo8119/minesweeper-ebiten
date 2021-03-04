package resource

import (
	"bytes"
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/gif"
	_ "image/png"
)

type _resources struct {
	Cursor1 *ebiten.Image
	Cursor2 *ebiten.Image

	Blank     *ebiten.Image
	Mousedown *ebiten.Image

	Face0 *ebiten.Image
	Face1 *ebiten.Image
	Face2 *ebiten.Image
	Face3 *ebiten.Image
	Face4 *ebiten.Image

	Num0 *ebiten.Image
	Num1 *ebiten.Image
	Num2 *ebiten.Image
	Num3 *ebiten.Image
	Num4 *ebiten.Image
	Num5 *ebiten.Image
	Num6 *ebiten.Image
	Num7 *ebiten.Image
	Num8 *ebiten.Image
	Num9 *ebiten.Image

	MineNum0 *ebiten.Image
	MineNum1 *ebiten.Image
	MineNum2 *ebiten.Image
	MineNum3 *ebiten.Image
	MineNum4 *ebiten.Image
	MineNum5 *ebiten.Image
	MineNum6 *ebiten.Image
	MineNum7 *ebiten.Image
	MineNum8 *ebiten.Image

	Mine0 *ebiten.Image
	Mine1 *ebiten.Image
	Mine2 *ebiten.Image

	Flag *ebiten.Image
}

var Resources = new(_resources)

func newImageFromFile(name string, fs embed.FS) *ebiten.Image {
	var (
		data []byte
		img  image.Image
		err  error
	)
	if data, err = fs.ReadFile(name); err != nil {
		panic(err)
	}
	if img, _, err = image.Decode(bytes.NewReader(data)); err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(img)
}

func Init(images embed.FS) {
	Resources.Cursor1 = newImageFromFile("images/cursor1.png", images)
	Resources.Cursor2 = newImageFromFile("images/cursor2.png", images)
	Resources.Blank = newImageFromFile("images/blank.gif", images)
	Resources.Mousedown = newImageFromFile("images/mousedown.gif", images)
	Resources.Face0 = newImageFromFile("images/face0.gif", images)
	Resources.Face1 = newImageFromFile("images/face1.gif", images)
	Resources.Face2 = newImageFromFile("images/face2.gif", images)
	Resources.Face3 = newImageFromFile("images/face3.gif", images)
	Resources.Face4 = newImageFromFile("images/face4.gif", images)
	Resources.Num0 = newImageFromFile("images/d0.gif", images)
	Resources.Num1 = newImageFromFile("images/d1.gif", images)
	Resources.Num2 = newImageFromFile("images/d2.gif", images)
	Resources.Num3 = newImageFromFile("images/d3.gif", images)
	Resources.Num4 = newImageFromFile("images/d4.gif", images)
	Resources.Num5 = newImageFromFile("images/d5.gif", images)
	Resources.Num6 = newImageFromFile("images/d6.gif", images)
	Resources.Num7 = newImageFromFile("images/d7.gif", images)
	Resources.Num8 = newImageFromFile("images/d8.gif", images)
	Resources.Num9 = newImageFromFile("images/d9.gif", images)
	Resources.MineNum0 = newImageFromFile("images/0.gif", images)
	Resources.MineNum1 = newImageFromFile("images/1.gif", images)
	Resources.MineNum2 = newImageFromFile("images/2.gif", images)
	Resources.MineNum3 = newImageFromFile("images/3.gif", images)
	Resources.MineNum4 = newImageFromFile("images/4.gif", images)
	Resources.MineNum5 = newImageFromFile("images/5.gif", images)
	Resources.MineNum6 = newImageFromFile("images/6.gif", images)
	Resources.MineNum7 = newImageFromFile("images/7.gif", images)
	Resources.MineNum8 = newImageFromFile("images/8.gif", images)
	Resources.Mine0 = newImageFromFile("images/mine0.gif", images)
	Resources.Mine1 = newImageFromFile("images/mine1.gif", images)
	Resources.Mine2 = newImageFromFile("images/mine2.gif", images)
	Resources.Flag = newImageFromFile("images/flag.gif", images)
}
