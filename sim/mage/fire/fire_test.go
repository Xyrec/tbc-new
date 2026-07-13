package fire

import (
	"testing"

	"github.com/wowsims/tbc/sim/common"
	_ "github.com/wowsims/tbc/sim/common"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

func init() {
	RegisterFireMage()
	common.RegisterAllEffects()
}

func TestFire(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class:      proto.Class_ClassMage,
			Race:       proto.Race_RaceTroll,
			OtherRaces: []proto.Race{proto.Race_RaceOrc},
			SpecOptions: core.SpecOptionsCombo{Label: "Fire", SpecOptions: &proto.Player_FireMage{
				FireMage: &proto.FireMage{
					Options: &proto.FireMage_Options{
						ClassOptions: &proto.MageOptions{
							DefaultMageArmor: proto.MageArmor_MageArmorMoltenArmor,
						},
					},
				},
			}},
			GearSet:  core.GetGearSet("../../../ui/mage/fire/gear_sets", "p2Fire"),
			Talents:  "20000000000000000000000-5052120123033310531251-0530020010000000000000",
			Rotation: core.GetAplRotation("../../../ui/mage/fire/apls", "fire"),
			ItemFilter: core.ItemFilter{
				WeaponTypes: []proto.WeaponType{
					proto.WeaponType_WeaponTypeDagger,
					proto.WeaponType_WeaponTypeSword,
					proto.WeaponType_WeaponTypeOffHand,
					proto.WeaponType_WeaponTypeStaff,
				},
				ArmorType: proto.ArmorType_ArmorTypeCloth,
				RangedWeaponTypes: []proto.RangedWeaponType{
					proto.RangedWeaponType_RangedWeaponTypeWand,
				},
				EnchantBlacklist: []int32{2673, 3225, 3273},
			},
		},
	}))
}
