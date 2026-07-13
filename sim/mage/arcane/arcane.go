package arcane

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/mage"
)

func RegisterArcaneMage() {
	core.RegisterAgentFactory(
		proto.Player_Mage{},
		proto.Spec_SpecMage,
		func(character *core.Character, options *proto.Player, _ *proto.Raid) core.Agent {
			return NewArcaneMage(character, options)
		},
		func(player *proto.Player, spec interface{}) {
			playerSpec, ok := spec.(*proto.Player_Mage)
			if !ok {
				panic("Invalid spec value for Arcane Mage!")
			}
			player.Spec = playerSpec
		},
	)
}

func NewArcaneMage(character *core.Character, options *proto.Player) *ArcaneMage {
	arcaneOptions := options.GetMage().Options

	return &ArcaneMage{
		Mage: mage.NewMage(character, options, arcaneOptions.ClassOptions),
	}
}

type ArcaneMage struct {
	*mage.Mage
}

func (arcaneMage *ArcaneMage) GetMage() *mage.Mage {
	return arcaneMage.Mage
}
