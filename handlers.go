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

	"github.com/EngoEngine/engo"
	"github.com/Noofbiz/engoBox2dSystem"
)

func addListeners(s *entityHolder) {
	engo.Mailbox.Listen("CollisionStartMessage", func(message engo.Message) {
		log.Printf("collision")
		c, isCollision := message.(engoBox2dSystem.CollisionStartMessage)
		if isCollision {
			a := c.Contact.GetFixtureA().GetBody().GetUserData()
			b := c.Contact.GetFixtureB().GetBody().GetUserData()
			if c.Contact.IsTouching() {
				for i1, e1 := range s.entities {
					log.Printf("i1: %d", i1)
					if e1.entity.BasicEntity.ID() == a || e1.entity.BasicEntity.ID() == b {
						log.Println("e1 obtained")
						for i2, e2 := range s.entities {
							log.Printf("i2: %d", i2)
							if e2.BasicEntity.ID() == a || e2.BasicEntity.ID() == b {
								// This means samus has hit the floor and is no longer jumping.
								// Stop the jump
								log.Println("e1 and e2 obtained")
								log.Println("jumpity jump")
								s.msys.samus.canJump = true
								s.msys.samus.totalJump = 0
							}
						}
					}
				}
			}
		}
	})
}
