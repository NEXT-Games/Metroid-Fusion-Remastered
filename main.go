package main

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"log"
)

type menuScene struct {}

type Samus struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func (*menuScene) Type() string { return "menuScene" }

func (*menuScene) Preload() {
	engo.Files.Load("tex/missingtex.jpg")
}

func (*menuScene) Setup(u engo.Updater) {
	// Setup Scene
	world, _ := u.(*ecs.World)
	world.AddSystem(&common.RenderSystem{})
	// Setup Samus
	sammy := Samus{BasicEntity: ecs.NewBasic()}
	sammy.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{0, 0},
		Width: 1024,
		Height: 576,
	}
	tex, err := common.LoadedSprite("tex/missingtex.jpg")
	if err !=  nil {
		log.Println("[FATAL] Can't load sprite for Samus! Error: " + err.Error())
	}
	sammy.RenderComponent = common.RenderComponent{
		Drawable: tex,
		Scale: engo.Point{1, 1},
	}
	for _, system := range world.Systems(){
		switch sys := system.(type){
		case *common.RenderSystem:
			sys.Add(&sammy.BasicEntity, &sammy.RenderComponent, &sammy.SpaceComponent)
		}
	}
}

func main() {
	opts := engo.RunOptions{
		Title: "METROID™️ FUSION: REMASTERED",
		Width: 1920,
		Height: 1080,
	}
	engo.Run(opts, &menuScene{})
}