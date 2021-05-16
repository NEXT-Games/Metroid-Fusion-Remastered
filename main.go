package main

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type menuScene struct {}

type Samus struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func (*menuScene) Type() string { return "menuScene" }

func (*menuScene) Preload() {
	engo.Files.Load("assets/placeholder.jpg")
}

func (*menuScene) Setup(engo.Updater) {
	// Setup Scene
	world, _ := u.(*ecs.World)
	world.AddSystem(&common.RenderSystem)
	// Setup Samus
	sammy := Samus(BasicEntity: ecs.NewBasic())
	sammy.SpaceComponent = common.SpaceComponent {
		Position: engo.Point{0, 0},
		Width: 1024,
		Height: 576,
	}
	tex, err := common.LoadedSprite("assets/placeholder.jpg")
	if err !=  nil {
		log.Println("[FATAL] Can't load sprite for Samus! Error: " + err.Error())
	}
	sammy.RenderComponent = common.RenderComponent{
		Drawable: tex,
		Scale: engo.point{1, 1}
	}
	for _, system := range world.Systems(){
		switch sys := system.(type)
		case *common.RenderSystems:
			sys.Add(&sammy.BasicEntity, &sammy.SpaceComponent, &sammy.RenderComponent)
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