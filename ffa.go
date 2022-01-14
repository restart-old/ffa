package ffa

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-plus/kit"
	"github.com/go-gl/mathgl/mgl64"
)

type FFA interface {
	Name() string
	World() *world.World
	Kit() kit.Kit
	Position() mgl64.Vec3
}

func TeleportToFFA(p *player.Player, ffa FFA) {
	ffa.World().AddEntity(p)
	p.Teleport(ffa.Position())

	p.Inventory().Clear()
	p.Armour().Clear()

	kit.GiveKit(p, ffa.Kit())
}
