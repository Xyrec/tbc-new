import * as PresetUtils from '../../core/preset_utils';
import {
	Debuffs,
	PseudoStat,
	RaidBuffs,
	Stat,
	ConsumesSpec,
	PartyBuffs,
	IndividualBuffs,
	TristateEffect,
	Profession,
	Drums,
	Spec,
} from '../../core/proto/common';
import { defaultImprovedShadowBoltSettings } from '../../core/proto_utils/utils';
import { Stats } from '../../core/proto_utils/stats';
import { SavedTalents } from '../../core/proto/ui';
import { MageArmor, FireMage_Options as FireMageOptions, FireMage_Rotation } from '../../core/proto/mage';
import BlankAPL from './apls/blank.apl.json';
import BlankGear from './gear_sets/blank.gear.json';
import FireApl from './apls/fire.apl.json';
import PreBISFireGear from './gear_sets/preBisFire.gear.json';
import P1FireGear from './gear_sets/p1Fire.gear.json';
import P2FireGear from './gear_sets/p2Fire.gear.json';
import { Phase } from '../../core/constants/other';
import { APLRotation_Type } from '../../core/proto/apl';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

export const BLANK_APL = PresetUtils.makePresetAPLRotation('Blank', BlankAPL);
export const PREBIS_FIRE = PresetUtils.makePresetGear('Fire PreRaid - BIS', PreBISFireGear, { phase: Phase.Phase1 });
export const P1_BIS_FIRE = PresetUtils.makePresetGear('Fire - BIS', P1FireGear, { phase: Phase.Phase1 });
export const P2_BIS_FIRE = PresetUtils.makePresetGear('Fire - BIS', P2FireGear, { phase: Phase.Phase2 });

export const FIRE_TALENTS = PresetUtils.makePresetTalents(
	'Fire',
	SavedTalents.create({ talentsString: '20000000000000000000000-5052120123033310531251-0530020010000000000000' }),
);
export const ROTATION_PRESET_FIRE = PresetUtils.makePresetAPLRotation('Fire', FireApl);
export const BLANK_GEARSET = PresetUtils.makePresetGear('Blank', BlankGear);

export const FireMageSimpleRotation = FireMage_Rotation.create({
	weaveFireBlast: true,
});

export const APL_FIRE_SIMPLE = PresetUtils.makePresetSimpleRotation('Fire', Spec.SpecFireMage, FireMageSimpleRotation);

// Preset options for EP weights
export const P1_EP_PRESET = PresetUtils.makePresetEpWeights(
	'P1 - Fire',
	Stats.fromMap(
		{
			[Stat.StatMana]: 0,
			[Stat.StatIntellect]: 0.22,
			[Stat.StatSpirit]: 0.11,
			[Stat.StatSpellDamage]: 1,
			[Stat.StatFireDamage]: 1,
			[Stat.StatFrostDamage]: 0,
			[Stat.StatSpellHitRating]: 1.7,
			[Stat.StatSpellCritRating]: 0.46,
			[Stat.StatSpellHasteRating]: 1.2,
			[Stat.StatSpellPenetration]: 0,
			[Stat.StatMP5]: 0.01,
		},
		{
			[PseudoStat.PseudoStatSchoolHitPercentFire]: 1.7,
		},
	),
);

export const P2_EP_PRESET = PresetUtils.makePresetEpWeights(
	'P2 - Fire',
	Stats.fromMap(
		{
			[Stat.StatMana]: 0,
			[Stat.StatIntellect]: 0.14,
			[Stat.StatSpirit]: 0.11,
			[Stat.StatSpellDamage]: 1,
			[Stat.StatFireDamage]: 1,
			[Stat.StatFrostDamage]: 0,
			[Stat.StatSpellHitRating]: 1.75,
			[Stat.StatSpellCritRating]: 0.46,
			[Stat.StatSpellHasteRating]: 1.35,
			[Stat.StatSpellPenetration]: 0,
			[Stat.StatMP5]: 0,
		},
		{
			[PseudoStat.PseudoStatSchoolHitPercentFire]: 1.71,
		},
	),
);

export const Talents = {
	name: 'Blank',
	data: SavedTalents.create({
		talentsString: '',
	}),
};

export const DefaultOptions = FireMageOptions.create({
	classOptions: {
		defaultMageArmor: MageArmor.MageArmorMoltenArmor,
	},
});

export const OtherDefaults = {
	distanceFromTarget: 20,
	profession1: Profession.Engineering,
	profession2: Profession.Tailoring,
};

export const DefaultConsumables = ConsumesSpec.create({
	guardianElixirId: 32067, // Elixir of Draenic Wisdom
	battleElixirId: 28103, // Adept's Elixir
	foodId: 27657, // Blackened Basilisk
	mhImbueId: 25122, // Brilliant Wizard Oil
	potId: 22839, // Destruction Potion
});

export const DefaultRaidBuffs = RaidBuffs.create({
	bloodlust: true,
	divineSpirit: 2,
	arcaneBrilliance: true,
	giftOfTheWild: 2,
	powerWordFortitude: 2,
	shadowProtection: true,
});

export const DefaultPartyBuffs = PartyBuffs.create({
	manaSpringTotem: 2,
	manaTideTotems: 1,
	wrathOfAirTotem: 1,
	totemOfWrath: 1,
	drums: Drums.LesserDrumsOfBattle,
});

export const DefaultIndividualBuffs = IndividualBuffs.create({
	blessingOfKings: true,
	blessingOfWisdom: 2,
	innervates: 1,
	powerInfusions: 1,
	shadowPriestDps: 1400,
});

export const DefaultDebuffs = Debuffs.create({
	misery: true,
	curseOfElements: 2,
	improvedSealOfTheCrusader: TristateEffect.TristateEffectImproved,
	judgementOfWisdom: true,
	...defaultImprovedShadowBoltSettings(),
});

export const P1_PLAYER_SETTINGS: PresetUtils.PresetSettings = {
	name: 'P1',
	playerOptions: OtherDefaults,
	debuffs: Debuffs.create({
		...DefaultDebuffs,
	}),
	reforgeSettings: {
		maxGemPhase: Phase.Phase1,
	},
};

export const P2_PLAYER_SETTINGS: PresetUtils.PresetSettings = {
	name: 'P2',
	playerOptions: OtherDefaults,
	partyBuffs: PartyBuffs.create({
		...DefaultPartyBuffs,
	}),
	debuffs: Debuffs.create({
		...DefaultDebuffs,
	}),
	reforgeSettings: {
		maxGemPhase: Phase.Phase2,
	},
};

export const P1_PRESET_BUILD_FIRE = PresetUtils.makePresetBuild('P1', {
	group: 'Fire',
	phase: Phase.Phase1,
	gear: P1_BIS_FIRE,
	talents: FIRE_TALENTS,
	epWeights: P1_EP_PRESET,
	rotationType: APLRotation_Type.TypeSimple,
	rotation: APL_FIRE_SIMPLE,
	settings: P1_PLAYER_SETTINGS,
});

export const P2_PRESET_BUILD_FIRE = PresetUtils.makePresetBuild('P2', {
	group: 'Fire',
	phase: Phase.Phase2,
	gear: P2_BIS_FIRE,
	talents: FIRE_TALENTS,
	epWeights: P2_EP_PRESET,
	rotationType: APLRotation_Type.TypeSimple,
	rotation: APL_FIRE_SIMPLE,
	settings: P2_PLAYER_SETTINGS,
});
