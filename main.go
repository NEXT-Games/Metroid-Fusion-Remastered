package main

import (
	"github.com/EngoEngine/engo"
)

type menuScene struct {}

func (*menuScene) Type() string { return "menuScene" }

func (*menuScene) Preload() {}

func (*menuScene) Setup(engo.Updater) {}

func main() {
	opts := engo.RunOptions{
		Title: "METROID™️ FUSION: REMASTERED",
		Width: 1920,
		Height: 1080,
	}
	engo.Run(opts, &menuScene{})
}