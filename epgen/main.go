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
	gearSet := os.Args[1]
	isbUptime, _ := strconv.ParseFloat(os.Args[2], 64)

	sim.RegisterAll()

	gear := core.GetGearSet("ui/mage/fire/gear_sets", gearSet)
	rotation := core.GetAplRotation("ui/mage/fire/apls", "fire")

	player := &proto.Player{
		Name:          "Fire Mage",
		Race:          proto.Race_RaceTroll,
		Class:         proto.Class_ClassMage,
		Equipment:     gear.GearSet,
		Rotation:      rotation.Rotation,
		TalentsString: "20000000000000000000000-5052120123033310531251-0530020010000000000000",
		Consumables: &proto.ConsumesSpec{
			GuardianElixirId: 32067,
			BattleElixirId:   28103,
			FoodId:           27657,
			MhImbueId:        25122,
			PotId:            22839,
		},
		Spec: &proto.Player_FireMage{
			FireMage: &proto.FireMage{
				Options: &proto.FireMage_Options{
					ClassOptions: &proto.MageOptions{
						DefaultMageArmor: proto.MageArmor_MageArmorMoltenArmor,
					},
				},
			},
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

	request := &proto.StatWeightsRequest{
		Player: player,
		RaidBuffs: &proto.RaidBuffs{
			Bloodlust:          true,
			DivineSpirit:       2,
			ArcaneBrilliance:   true,
			GiftOfTheWild:      2,
			PowerWordFortitude: 2,
			ShadowProtection:   true,
		},
		PartyBuffs: &proto.PartyBuffs{
			ManaSpringTotem: 2,
			ManaTideTotems:  1,
			WrathOfAirTotem: 1,
			TotemOfWrath:    1,
			Drums:           proto.Drums_LesserDrumsOfBattle,
		},
		Debuffs: &proto.Debuffs{
			Misery:                    true,
			CurseOfElements:           2,
			ImprovedSealOfTheCrusader: proto.TristateEffect_TristateEffectImproved,
			JudgementOfWisdom:         true,
			IsbUptime:                 isbUptime,
		},
		Encounter:  core.MakeSingleTargetEncounter(0.1),
		SimOptions: &proto.SimOptions{Iterations: 30000, RandomSeed: 101},
		StatsToWeigh: []proto.Stat{
			proto.Stat_StatIntellect,
			proto.Stat_StatSpirit,
			proto.Stat_StatSpellDamage,
			proto.Stat_StatArcaneDamage,
			proto.Stat_StatFireDamage,
			proto.Stat_StatFrostDamage,
			proto.Stat_StatSpellHitRating,
			proto.Stat_StatSpellCritRating,
			proto.Stat_StatSpellHasteRating,
			proto.Stat_StatMana,
			proto.Stat_StatMP5,
		},
		PseudoStatsToWeigh: []proto.PseudoStat{
			proto.PseudoStat_PseudoStatSchoolHitPercentArcane,
			proto.PseudoStat_PseudoStatSchoolHitPercentFire,
			proto.PseudoStat_PseudoStatSchoolHitPercentFrost,
		},
		EpReferenceStat: proto.Stat_StatSpellDamage,
	}

	result := core.StatWeights(request)
	if result.Error != nil {
		fmt.Println("ERROR:", result.Error.Message)
		os.Exit(1)
	}

	fmt.Printf("=== %s (isbUptime %.2f) ===\n", gearSet, isbUptime)
	for _, s := range request.StatsToWeigh {
		fmt.Printf("stat %-22s ep %.3f (stdev %.3f)\n", s.String(), result.Dps.EpValues.Stats[s], result.Dps.EpValuesStdev.Stats[s])
	}
	for _, ps := range request.PseudoStatsToWeigh {
		fmt.Printf("pseudo %-20s ep %.3f (stdev %.3f)\n", ps.String(), result.Dps.EpValues.PseudoStats[ps], result.Dps.EpValuesStdev.PseudoStats[ps])
	}
}
