package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/wowsims/tbc/sim"
	"github.com/wowsims/tbc/sim/core"
	"github.com/wowsims/tbc/sim/core/proto"
)

func main() {
	spec := os.Args[1]
	gearSet := os.Args[2]
	isbUptime, _ := strconv.ParseFloat(os.Args[3], 64)
	durationSec, _ := strconv.Atoi(os.Args[4])

	sim.RegisterAll()

	gear := core.GetGearSet("ui/mage/"+spec+"/gear_sets", gearSet)
	rotation := core.GetAplRotation("ui/mage/"+spec+"/apls", spec)

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
	}

	player := &proto.Player{
		Name:          "Mage",
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
	applySpec(player)

	encounter := core.MakeSingleTargetEncounter(0.05)
	encounter.Duration = float64(time.Duration(durationSec) * time.Second / time.Second)

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
				IsbUptime:                 isbUptime,
			},
		},
		Encounter:  encounter,
		SimOptions: &proto.SimOptions{Iterations: 10000, RandomSeed: 101, Debug: os.Getenv("SIMDEBUG") != "", DebugFirstIteration: os.Getenv("SIMDEBUG") != ""},
	}

	if os.Getenv("SIMDEBUG") != "" {
		rsr.SimOptions.Iterations = 1
	}
	result := core.RunRaidSim(rsr)
	if os.Getenv("SIMDEBUG") != "" {
		fmt.Print(result.Logs)
		return
	}
	fmt.Printf("dps: %.1f (duration %ds)\n", result.RaidMetrics.Dps.Avg, durationSec)

	type row struct {
		name               string
		dmg                float64
		casts, hits, crits int32
	}
	var rows []row
	var total float64
	collect := func(prefix string, actions []*proto.ActionMetrics) {
		for _, a := range actions {
			var dmg float64
			var casts, hits, crits int32
			for _, t := range a.Targets {
				dmg += t.Damage
				casts += t.Casts
				hits += t.Hits
				crits += t.Crits
			}
			total += dmg
			rows = append(rows, row{prefix + a.Id.String(), dmg, casts, hits, crits})
		}
	}
	for _, p := range result.RaidMetrics.Parties[0].Players {
		collect(p.Name+" ", p.Actions)
		for _, pet := range p.Pets {
			collect(pet.Name+" ", pet.Actions)
		}
	}
	sort.Slice(rows, func(i, j int) bool { return rows[i].dmg > rows[j].dmg })
	for _, r := range rows {
		if r.dmg == 0 {
			continue
		}
		critPct := 0.0
		if r.hits > 0 {
			critPct = 100 * float64(r.crits) / float64(r.hits)
		}
		fmt.Printf("%-70s %6.2f%%  casts/iter %5.1f  crit %5.1f%%\n", r.name, 100*r.dmg/total, float64(r.casts)/10000, critPct)
	}
}
