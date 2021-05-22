package main

import (
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type menuScene struct{}

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
	engo.Input.RegisterButton("MoveLeft", engo.KeyA)
	engo.Input.RegisterButton("MoveRight", engo.KeyD)
	engo.Input.RegisterButton("Jump", engo.KeySpace)
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&movingThingSystem{})
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
		case *movingThingSystem:
			sys.Add(&sammy.BasicEntity, &sammy.RenderComponent, &sammy.SpaceComponent)
		}
	}
	log.Println("Designed with ❤️ by NEXT Games")
}

type movingThingSystem struct {
	spaceComponent *common.SpaceComponent
	totalJump      int
	isJumping      bool
}

func (*movingThingSystem) Type() string { return "movingThingSystem" }

func (self *movingThingSystem) Update(dt float32) {
	// self.spaceComponent.Position = engo.Point{100, 100}
	// A friendly reminder that **we do NOT do a little trolling**
	if engo.Input.Button("MoveLeft").Down() {
		self.spaceComponent.Position.X -= 3
	}
	if engo.Input.Button("MoveRight").Down() {
		self.spaceComponent.Position.X += 3
	}
	if engo.Input.Button("Jump").Down() && !self.isJumping && self.totalJump <= 100 {
		self.spaceComponent.Position.Y -= 10
		self.totalJump += 10
		self.isJumping = true
	}
	if self.isJumping && self.totalJump <= 100 {
		self.spaceComponent.Position.Y -= 10
		self.totalJump += 10
	}
	if self.isJumping && self.totalJump >= 100 {
		self.totalJump = 0
		self.isJumping = false
	}
	if self.totalJump < 0 {
		// todo
	}
}

func (self *movingThingSystem) Add(basicEntity *ecs.BasicEntity, renderComponent *common.RenderComponent, spaceComponent *common.SpaceComponent) {
	self.spaceComponent = spaceComponent
	self.isJumping = false
	self.totalJump = 0
}

func (self *movingThingSystem) Remove(added ecs.BasicEntity) {
	// nop
}

func main() {
	opts := engo.RunOptions{
		Title:  "METROID™️ FUSION: REMASTERED",
		Width:  1920,
		Height: 1080,
	}
	engo.Run(opts, &menuScene{})
}
