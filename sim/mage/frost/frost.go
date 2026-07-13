package frost

import (
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	"github.com/wowsims/tbc/sim/mage"
)

func RegisterFrostMage() {
	core.RegisterAgentFactory(
		proto.Player_FrostMage{},
		proto.Spec_SpecFrostMage,
		func(character *core.Character, options *proto.Player, _ *proto.Raid) core.Agent {
			return NewFrostMage(character, options)
		},
		func(player *proto.Player, spec interface{}) {
			playerSpec, ok := spec.(*proto.Player_FrostMage)
			if !ok {
				panic("Invalid spec value for Frost Mage!")
			}
			player.Spec = playerSpec
		},
	)
}

func NewFrostMage(character *core.Character, options *proto.Player) *FrostMage {
	frostOptions := options.GetFrostMage().Options

	return &FrostMage{
		Mage: mage.NewMage(character, options, frostOptions.ClassOptions),
	}
}

type FrostMage struct {
	*mage.Mage
}

func (frostMage *FrostMage) GetMage() *mage.Mage {
	return frostMage.Mage
}
