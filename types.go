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

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
	"github.com/Noofbiz/engoBox2dSystem"
)

type BaseEntity struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	engoBox2dSystem.Box2dComponent
	spriteMeta string
	totalJump  int
	canJump    bool
}
type entityType struct {
	ecs.BasicEntity
	*engoBox2dSystem.Box2dComponent
	entity BaseEntity
}

type entityHolder struct {
	entities []*entityType
	msys     *movementSystem
}

type DummyMessage struct{}

type Text struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.RenderComponent
	text string
	font common.Font
}

func (entity *entityType) Debug() {
	log.Printf("debug entityType totaljump: %d", entity.entity.totalJump)
	if entity.entity.canJump {
		log.Println("debug entityType success canJump")
	}
}
func (holder *entityHolder) Add(e *entityType) {
	holder.entities = append(holder.entities, e)
}
func (holder *entityHolder) SetMsys(sys *movementSystem) {
	holder.msys = sys
}
func (holder *entityType) SetCanJump() {
	holder.entity.canJump = true
	holder.entity.totalJump = 0
}

func (*DummyMessage) Type() string {
	return "menuswitch"
}
