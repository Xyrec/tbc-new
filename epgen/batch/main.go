package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/wowsims/tbc/sim"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

type combo struct {
	Name      string          `json:"name"`
	Equipment json.RawMessage `json:"equipment"`
}

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
	combosFile := os.Args[2]

	sim.RegisterAll()

	data, err := os.ReadFile(combosFile)
	if err != nil {
		panic(err)
	}
	var combos []combo
	if err := json.Unmarshal(data, &combos); err != nil {
		panic(err)
	}

	rotation := core.GetAplRotation("ui/mage/"+spec+"/apls", spec)

	type result struct {
		name string
		dps  float64
	}
	var results []result

	for _, c := range combos {
		equipment := &proto.EquipmentSpec{}
		if err := protojson.Unmarshal(c.Equipment, equipment); err != nil {
			panic(c.Name + ": " + err.Error())
		}

		player := &proto.Player{
			Name:          "Fire Mage",
			Race:          proto.Race_RaceTroll,
			Class:         proto.Class_ClassMage,
			Equipment:     equipment,
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

		applySpec(player)
		rsr := &proto.RaidSimRequest{
			Raid: &proto.Raid{
				Parties: []*proto.Party{{
					Players: []*proto.Player{player},
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
					IsbUptime:                 0.59,
				},
			},
			Encounter:  core.MakeSingleTargetEncounter(0.1),
			SimOptions: &proto.SimOptions{Iterations: 30000, RandomSeed: 101},
		}

		dps := core.RunRaidSim(rsr).RaidMetrics.Dps.Avg
		results = append(results, result{c.Name, dps})
		fmt.Printf("%-32s %.2f dps\n", c.Name, dps)
	}

	var baseline float64
	for _, r := range results {
		if r.name == "A-current(T5-2pc,no-others)" {
			baseline = r.dps
		}
	}
	sort.Slice(results, func(i, j int) bool { return results[i].dps > results[j].dps })
	fmt.Println("\n=== RANKED ===")
	for i, r := range results {
		fmt.Printf("%d. %-32s %.2f (%+.1f vs current)\n", i+1, r.name, r.dps, r.dps-baseline)
	}
}
