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
	"log"

	"github.com/ByteArena/box2d"
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type movementSystem struct {
	spaceComponent *common.SpaceComponent
	samus          BaseEntity
}

func (*movementSystem) Type() string { return "movementSystem" }

func (movementSystem *movementSystem) Update(dt float32) {
	// movementSystem.spaceComponent.Position = engo.Point{100, 100}
	// A friendly reminder that **we do NOT do a little trolling**
	if engo.Input.Button("MoveLeft").Down() {
		movementSystem.spaceComponent.Position.X -= 3
	}
	if engo.Input.Button("MoveRight").Down() {
		movementSystem.spaceComponent.Position.X += 3
	}
	if engo.Input.Button("Jump").Down() && movementSystem.samus.totalJump < 100 {
		movementSystem.samus.Body.ApplyLinearImpulseToCenter(box2d.B2Vec2{X: 0, Y: -2000}, true)
		movementSystem.samus.totalJump += 5
	}
	if engo.Input.Button("Jump").Down() && !movementSystem.samus.canJump {
		log.Printf("%d", movementSystem.samus.totalJump)
	}
}
func (movementSystem *movementSystem) AddEtc(samus BaseEntity) {
	movementSystem.samus = samus
	movementSystem.samus.canJump = true
}
func (movementSystem *movementSystem) Add(basicEntity *ecs.BasicEntity, renderComponent *common.RenderComponent, spaceComponent *common.SpaceComponent) {
	movementSystem.spaceComponent = spaceComponent
}

func (movementSystem *movementSystem) Remove(added ecs.BasicEntity) {
	// nop
}

type menuSystem struct{}

func (*menuSystem) Type() string { return "menuSystem" }
func (sys menuSystem) Update(dt float32) {
	if engo.Input.Button("startgame").JustPressed() {
		engo.Mailbox.Dispatch(&DummyMessage{})
		engo.SetScene(&MainDeckScene{}, true)
	}
}
func (sys menuSystem) Add(e *ecs.BasicEntity) {
	// nop
}

func (sys menuSystem) Remove(e ecs.BasicEntity) {
	// nop
}
