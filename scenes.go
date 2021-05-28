package main

import (
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type MainDeckScene struct{}

type Samus struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func (*MainDeckScene) Type() string { return "MainDeckScene" }

func (*MainDeckScene) Preload() {
	engo.Files.Load("tex/missingtex.jpg")
}

func (*MainDeckScene) Setup(u engo.Updater) {
	// Setup Scene
	world, _ := u.(*ecs.World)
	engo.Input.RegisterButton("MoveLeft", engo.KeyA)
	engo.Input.RegisterButton("MoveRight", engo.KeyD)
	engo.Input.RegisterButton("Jump", engo.KeySpace)
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&movementSystem{})
	// Setup Samus
	sammy := Samus{BasicEntity: ecs.NewBasic()}
	sammy.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{0, 0},
		Width:    1024,
		Height:   576,
	}
	tex, err := common.LoadedSprite("tex/missingtex.jpg")
	if err != nil {
		log.Println("[FATAL] Can't load sprite for Samus! Error: " + err.Error())
	}
	sammy.RenderComponent = common.RenderComponent{
		Drawable: tex,
		Scale:    engo.Point{1, 1},
	}
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&sammy.BasicEntity, &sammy.RenderComponent, &sammy.SpaceComponent)
		case *movementSystem:
			sys.Add(&sammy.BasicEntity, &sammy.RenderComponent, &sammy.SpaceComponent)
		}
	}
	log.Println("Designed with ❤️ by NEXT Games")
	log.Println("If you have paid for this software you have been scammed")
}
