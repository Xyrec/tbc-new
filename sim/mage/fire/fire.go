package fire

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/mage"
)

func RegisterFireMage() {
	core.RegisterAgentFactory(
		proto.Player_FireMage{},
		proto.Spec_SpecFireMage,
		func(character *core.Character, options *proto.Player, _ *proto.Raid) core.Agent {
			return NewFireMage(character, options)
		},
		func(player *proto.Player, spec interface{}) {
			playerSpec, ok := spec.(*proto.Player_FireMage)
			if !ok {
				panic("Invalid spec value for Fire Mage!")
			}
			player.Spec = playerSpec
		},
	)
}

func NewFireMage(character *core.Character, options *proto.Player) *FireMage {
	fireOptions := options.GetFireMage().Options

	return &FireMage{
		Mage: mage.NewMage(character, options, fireOptions.ClassOptions),
	}
}

type FireMage struct {
	*mage.Mage
}

func (fireMage *FireMage) GetMage() *mage.Mage {
	return fireMage.Mage
}
