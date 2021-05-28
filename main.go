package main

import (
	"github.com/EngoEngine/engo"
)

func main() {
	opts := engo.RunOptions{
		Title:  "METROID™️ FUSION: REMASTERED",
		Width:  1920,
		Height: 1080,
	}
	engo.Run(opts, &MainDeckScene{})
}
