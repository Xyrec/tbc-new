package frost

import (
	"testing"

	"github.com/wowsims/tbc/sim/common"
	_ "github.com/wowsims/tbc/sim/common"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

func init() {
	RegisterFrostMage()
	common.RegisterAllEffects()
}

func TestFrost(t *testing.T) {
	core.RunTestSuite(t, t.Name(), core.FullCharacterTestSuiteGenerator([]core.CharacterSuiteConfig{
		{
			Class:      proto.Class_ClassMage,
			Race:       proto.Race_RaceTroll,
			OtherRaces: []proto.Race{proto.Race_RaceOrc},
			SpecOptions: core.SpecOptionsCombo{Label: "Frost", SpecOptions: &proto.Player_FrostMage{
				FrostMage: &proto.FrostMage{
					Options: &proto.FrostMage_Options{
						ClassOptions: &proto.MageOptions{
							DefaultMageArmor: proto.MageArmor_MageArmorMoltenArmor,
						},
					},
				},
			}},
			GearSet:  core.GetGearSet("../../../ui/mage/frost/gear_sets", "p2Frost"),
			Talents:  "230005--0535020310235310250551",
			Rotation: core.GetAplRotation("../../../ui/mage/frost/apls", "frost"),
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
