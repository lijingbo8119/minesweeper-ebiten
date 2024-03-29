package main

import (
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lijingbo8119/minesweeper-ebiten/game"
	"github.com/lijingbo8119/minesweeper-ebiten/resource"
	_ "image/png"
	"log"
)

//go:embed images/*
var images embed.FS

var gameInstance = game.NewGame(game.MainBoard{})

func main() {
	resource.Init(images)
	ebiten.SetWindowSize(game.ScreenWidth*2, game.ScreenHeight*2)
	ebiten.SetWindowTitle("Minesweeper")
	if err := ebiten.RunGame(gameInstance); err != nil {
		log.Fatal(err)
	}
}
