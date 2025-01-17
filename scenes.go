/*
Copyright (c) 2021, Warp Studios
All rights reserved.

Metroid is (C) Nintendo
All rights reserved.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package main

import (
	"bytes"
	"image/color"
	"log"

	"github.com/ByteArena/box2d"
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/Noofbiz/engoBox2dSystem"
	"golang.org/x/image/font/gofont/gosmallcaps"
)

type MainDeckScene struct{}

func (*MainDeckScene) Type() string { return "MainDeckScene" }

func (*MainDeckScene) Preload() {
	engo.Files.Load("tex/missingtex.jpg")
}

func (*MainDeckScene) Setup(u engo.Updater) {
	// Setup Scene
	world, _ := u.(*ecs.World)
	engo.Mailbox.Dispatch(&DummyMessage{})
	engo.Input.RegisterButton("MoveLeft", engo.KeyA)
	engo.Input.RegisterButton("MoveRight", engo.KeyD)
	engo.Input.RegisterButton("Jump", engo.KeySpace)
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&movementSystem{})
	world.AddSystem(&engoBox2dSystem.PhysicsSystem{VelocityIterations: 3, PositionIterations: 8})
	world.AddSystem(&engoBox2dSystem.CollisionSystem{})

	engoBox2dSystem.World.SetGravity(box2d.B2Vec2{X: 0, Y: 10})
	// Setup Samus
	sammy := BaseEntity{BasicEntity: ecs.NewBasic(), spriteMeta: "samus"}
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
	sammyDef := box2d.NewB2BodyDef()
	sammyDef.Type = box2d.B2BodyType.B2_dynamicBody
	sammyDef.Position = engoBox2dSystem.Conv.ToBox2d2Vec(sammy.Center())
	sammyDef.Angle = engoBox2dSystem.Conv.DegToRad(sammy.Rotation)
	sammy.Box2dComponent.Body = engoBox2dSystem.World.CreateBody(sammyDef)
	var sammyBodyShape box2d.B2PolygonShape

	sammyBodyShape.SetAsBox(engoBox2dSystem.Conv.PxToMeters(sammy.SpaceComponent.Width/2), engoBox2dSystem.Conv.PxToMeters(sammy.SpaceComponent.Height/2))
	sammyFixtureDef := box2d.B2FixtureDef{
		Shape:    &sammyBodyShape,
		Density:  1.0,
		Friction: 0.1,
	}
	sammy.Box2dComponent.Body.CreateFixtureFromDef(&sammyFixtureDef)
	testGround := BaseEntity{BasicEntity: ecs.NewBasic(), spriteMeta: "ground"}
	testGround.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{0, 1000},
		Width:    1024,
		Height:   576,
	}
	tex, err = common.LoadedSprite("tex/missingtex.jpg")
	if err != nil {
		log.Println("[FATAL] Can't load sprite for Samus! Error: " + err.Error())
	}
	testGround.RenderComponent = common.RenderComponent{
		Drawable: tex,
		Scale:    engo.Point{1, 1},
	}
	grassBodyDef := box2d.NewB2BodyDef()
	grassBodyDef.Position = engoBox2dSystem.Conv.ToBox2d2Vec(testGround.Center())
	grassBodyDef.Angle = engoBox2dSystem.Conv.DegToRad(testGround.Rotation)
	testGround.Box2dComponent.Body = engoBox2dSystem.World.CreateBody(grassBodyDef)
	var grassBodyShape box2d.B2PolygonShape
	grassBodyShape.SetAsBox(engoBox2dSystem.Conv.PxToMeters(testGround.SpaceComponent.Width/2),
		engoBox2dSystem.Conv.PxToMeters(testGround.SpaceComponent.Height/2))
	grassFixtureDef := box2d.B2FixtureDef{Shape: &grassBodyShape}
	testGround.Box2dComponent.Body.CreateFixtureFromDef(&grassFixtureDef)
	entityholder := entityHolder{}
	entityholder.Add(&entityType{sammy.BasicEntity, &sammy.Box2dComponent, sammy})
	entityholder.Add(&entityType{testGround.BasicEntity, &testGround.Box2dComponent, testGround})
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&sammy.BasicEntity, &sammy.RenderComponent, &sammy.SpaceComponent)
			sys.Add(&testGround.BasicEntity, &testGround.RenderComponent, &testGround.SpaceComponent)
		case *movementSystem:
			sys.Add(&sammy.BasicEntity, &sammy.RenderComponent, &sammy.SpaceComponent)
			sys.AddEtc(sammy)
			entityholder.SetMsys(sys)
		case *engoBox2dSystem.PhysicsSystem:
			sys.Add(&sammy.BasicEntity, &sammy.SpaceComponent, &sammy.Box2dComponent)
			sys.Add(&testGround.BasicEntity, &testGround.SpaceComponent, &testGround.Box2dComponent)
		case *engoBox2dSystem.CollisionSystem:
			sys.Add(&sammy.BasicEntity, &sammy.SpaceComponent, &sammy.Box2dComponent)
			sys.Add(&testGround.BasicEntity, &testGround.SpaceComponent, &testGround.Box2dComponent)
		}
	}

	addListeners(&entityholder)
	log.Println("Designed with ❤️ by NEXT Games")
	log.Println("If you have paid for this software you have been scammed")
}

type MenuScene struct{}

type Music struct {
	ecs.BasicEntity
	common.AudioComponent
}

func (*MenuScene) Preload() {
	engo.Files.Load("audio/theme_of_m4r.mp3")
	engo.Files.LoadReaderData("go.ttf", bytes.NewReader(gosmallcaps.TTF))
}

func (scene *MenuScene) Setup(u engo.Updater) {
	world := u.(*ecs.World)
	as := &common.AudioSystem{}
	world.AddSystem(as)
	world.AddSystem(menuSystem{})
	world.AddSystem(&common.RenderSystem{})
	menutheme := Music{BasicEntity: ecs.NewBasic()}
	var err error
	menutheme.AudioComponent.Player, err = common.LoadedPlayer("audio/theme_of_m4r.mp3")
	if err != nil {
		log.Fatal(err)
	}
	as.Add(&menutheme.BasicEntity, &menutheme.AudioComponent)
	menutheme.Player.Play()
	menutheme.Player.Repeat = true
	fnt := &common.Font{
		URL:  "go.ttf",
		FG:   color.White,
		Size: 24,
	}
	fnt.CreatePreloaded()
	text := Text{BasicEntity: ecs.NewBasic(), text: "Press Enter", font: fnt}
	text.RenderComponent.Drawable = common.Text{Font: fnt, Text: "Press Enter"}
	bogus := Text{BasicEntity: ecs.NewBasic(), text: "sus imposter", font: fnt}
	bogus.RenderComponent.Drawable = common.Text{Font: fnt, Text: "sus imposter"}
	text.SetShader(common.TextHUDShader)
	text.RenderComponent.SetZIndex(1001)
	text.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 0, Y: 0},
		Width:    1000,
		Height:   1000,
	}
	bogus.SetShader(common.TextHUDShader)
	bogus.RenderComponent.SetZIndex(1001)
	bogus.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 0, Y: 100},
		Width:    1000,
		Height:   1000,
	}
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&text.BasicEntity, &text.RenderComponent, &text.SpaceComponent)
			sys.Add(&bogus.BasicEntity, &bogus.RenderComponent, &bogus.SpaceComponent)
		case *menuSystem:
			sys.Add(text)
			sys.Add(bogus)
		}

	}
	engo.Input.RegisterButton("startgame", engo.KeyEnter)
	engo.Mailbox.Listen("menuswitch", func(message engo.Message) {
		log.Println("menuswitch")
		menutheme.Player.Close()
	})
}

func (scene *MenuScene) Type() string {
	return "MenuScene"
}
