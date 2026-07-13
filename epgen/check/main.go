package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/wowsims/tbc/sim"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

func main() {
	spec := os.Args[1]
	var talents string
	var applySpec func(p *proto.Player)
	switch spec {
	case "fire":
		talents = "20000000000000000000000-5052120123033310531251-0530020010000000000000"
		applySpec = func(p *proto.Player) {
			p.Spec = &proto.Player_FireMage{
				FireMage: &proto.FireMage{
					Options: &proto.FireMage_Options{
						ClassOptions: &proto.MageOptions{DefaultMageArmor: proto.MageArmor_MageArmorMoltenArmor},
					},
				},
			}
		}
	case "arcane":
		talents = "2500052300030150330125--053500031003001"
		applySpec = func(p *proto.Player) {
			p.Spec = &proto.Player_Mage{
				Mage: &proto.Mage{
					Options: &proto.Mage_Options{
						ClassOptions: &proto.MageOptions{DefaultMageArmor: proto.MageArmor_MageArmorMageArmor},
					},
				},
			}
		}
	case "frost":
		talents = "230005--0535020310235310250551"
		applySpec = func(p *proto.Player) {
			p.Spec = &proto.Player_FrostMage{
				FrostMage: &proto.FrostMage{
					Options: &proto.FrostMage_Options{
						ClassOptions: &proto.MageOptions{DefaultMageArmor: proto.MageArmor_MageArmorMoltenArmor},
					},
				},
			}
		}
	default:
		panic("unknown spec: " + spec)
	}
	gearSet := os.Args[2]
	isbUptime, _ := strconv.ParseFloat(os.Args[3], 64)

	sim.RegisterAll()

	gear := core.GetGearSet("ui/mage/"+spec+"/gear_sets", gearSet)
	rotation := core.GetAplRotation("ui/mage/"+spec+"/apls", spec)

	makePlayer := func(stat int, bonus float64) *proto.Player {
		p := &proto.Player{
			Name:          "Fire Mage",
			Race:          proto.Race_RaceTroll,
			Class:         proto.Class_ClassMage,
			Equipment:     gear.GearSet,
			Rotation:      rotation.Rotation,
			TalentsString: talents,
			Consumables: &proto.ConsumesSpec{
				GuardianElixirId: 32067,
				BattleElixirId:   28103,
				FoodId:           27657,
				MhImbueId:        25122,
				PotId:            22839,
			},
			Buffs: &proto.IndividualBuffs{
				BlessingOfKings:  true,
				BlessingOfWisdom: 2,
				Innervates:       1,
				PowerInfusions:   1,
				ShadowPriestDps:  1400,
			},
			Profession1:        proto.Profession_Engineering,
			Profession2:        proto.Profession_Tailoring,
			DistanceFromTarget: 20,
			ReactionTimeMs:     100,
		}
		applySpec(p)
		if bonus != 0 {
			if stat >= 1000 {
				pseudo := make([]float64, 30)
				pseudo[stat-1000] = bonus
				p.BonusStats = &proto.UnitStats{PseudoStats: pseudo}
			} else {
				stats := make([]float64, 42)
				stats[stat] = bonus
				p.BonusStats = &proto.UnitStats{Stats: stats}
			}
		}
		return p
	}

	run := func(stat int, bonus float64) float64 {
		rsr := &proto.RaidSimRequest{
			Raid: &proto.Raid{
				Parties: []*proto.Party{{
					Players: []*proto.Player{makePlayer(stat, bonus)},
					Buffs: &proto.PartyBuffs{
						ManaSpringTotem: 2,
						ManaTideTotems:  1,
						WrathOfAirTotem: 1,
						TotemOfWrath:    1,
						Drums:           proto.Drums_LesserDrumsOfBattle,
					},
				}},
				Buffs: &proto.RaidBuffs{
					Bloodlust:          true,
					DivineSpirit:       2,
					ArcaneBrilliance:   true,
					GiftOfTheWild:      2,
					PowerWordFortitude: 2,
					ShadowProtection:   true,
				},
				Debuffs: &proto.Debuffs{
					Misery:                    true,
					CurseOfElements:           2,
					ImprovedSealOfTheCrusader: proto.TristateEffect_TristateEffectImproved,
					JudgementOfWisdom:         true,
					IsbUptime:                 isbUptime,
				},
			},
			Encounter:  core.MakeSingleTargetEncounter(0.1),
			SimOptions: &proto.SimOptions{Iterations: 30000, RandomSeed: 101},
		}
		return core.RunRaidSim(rsr).RaidMetrics.Dps.Avg
	}

	type sweep struct {
		name  string
		stat  int
		delta float64
	}
	sweeps := []sweep{
		{"Intellect", 3, 100},
		{"Spirit", 16, 100},
		{"SpellDamage", 5, 100},
		{"SpellHitRating", 12, -50},
		{"SpellCritRating", 13, 100},
		{"SpellHasteRating", 14, 100},
		{"Mana", 34, 500},
		{"MP5", 35, 25},
	}
	if spec == "fire" {
		sweeps = append(sweeps, sweep{"FireDamage", 7, 100}, sweep{"SchoolHitPctFire", 1007, -2})
	} else if spec == "arcane" {
		sweeps = append(sweeps, sweep{"ArcaneDamage", 6, 100}, sweep{"SchoolHitPctArcane", 1006, -2})
	} else {
		sweeps = append(sweeps, sweep{"FrostDamage", 8, 100}, sweep{"SchoolHitPctFrost", 1008, -2})
	}
	if len(os.Args) > 4 {
		stat, _ := strconv.Atoi(os.Args[4])
		delta, _ := strconv.ParseFloat(os.Args[5], 64)
		sweeps = []sweep{{"SpellDamage", 5, 100}, {"Custom", stat, delta}}
	}

	base := run(0, 0)
	fmt.Printf("=== %s (isb %.2f) base dps %.2f ===\n", gearSet, isbUptime, base)
	var ref float64
	for _, s := range sweeps {
		d := run(s.stat, s.delta)
		perPoint := (d - base) / s.delta
		if s.name == "SpellDamage" {
			ref = perPoint
		}
		fmt.Printf("%-18s delta %+5.0f dps/pt %+.4f\n", s.name, s.delta, perPoint)
	}
	fmt.Printf("reference (SpellDamage) = %.4f dps/pt; divide others by this for EP\n", ref)
}
