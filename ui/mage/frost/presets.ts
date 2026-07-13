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
} from '../../core/proto/common';
import { defaultImprovedShadowBoltSettings } from '../../core/proto_utils/utils';
import { Stats } from '../../core/proto_utils/stats';
import { SavedTalents } from '../../core/proto/ui';
import { MageArmor, FrostMage_Options as FrostMageOptions } from '../../core/proto/mage';
import BlankAPL from './apls/blank.apl.json';
import BlankGear from './gear_sets/blank.gear.json';
import FrostApl from './apls/frost.apl.json';
import PreBISFrostGear from './gear_sets/preBisFrost.gear.json';
import P1FrostGear from './gear_sets/p1Frost.gear.json';
import P2FrostGear from './gear_sets/p2Frost.gear.json';
import { Phase } from '../../core/constants/other';
import { APLRotation_Type } from '../../core/proto/apl';

// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.

export const BLANK_APL = PresetUtils.makePresetAPLRotation('Blank', BlankAPL);
export const PREBIS_FROST = PresetUtils.makePresetGear('Frost PreRaid - BIS', PreBISFrostGear, { phase: Phase.Phase1 });
export const P1_BIS_FROST = PresetUtils.makePresetGear('Frost - BIS', P1FrostGear, { phase: Phase.Phase1 });
export const P2_BIS_FROST = PresetUtils.makePresetGear('Frost - BIS', P2FrostGear, { phase: Phase.Phase2 });

export const FROST_TALENTS = PresetUtils.makePresetTalents('Frost', SavedTalents.create({ talentsString: '230005--0535020310235310250551' }));
export const ROTATION_PRESET_FROST = PresetUtils.makePresetAPLRotation('Frost', FrostApl);
export const BLANK_GEARSET = PresetUtils.makePresetGear('Blank', BlankGear);

// Preset options for EP weights
export const P1_EP_PRESET = PresetUtils.makePresetEpWeights(
	'P1 - Frost',
	Stats.fromMap(
		{
			[Stat.StatMana]: 0,
			[Stat.StatIntellect]: 0.23,
			[Stat.StatSpirit]: 0.11,
			[Stat.StatSpellDamage]: 1,
			[Stat.StatFrostDamage]: 1,
			[Stat.StatFireDamage]: 0,
			[Stat.StatSpellHitRating]: 1.57,
			[Stat.StatSpellCritRating]: 0.75,
			[Stat.StatSpellHasteRating]: 1.18,
			[Stat.StatSpellPenetration]: 0,
			[Stat.StatMP5]: 0.01,
		},
		{
			[PseudoStat.PseudoStatSchoolHitPercentFrost]: 1.57,
		},
	),
);

export const P2_EP_PRESET = PresetUtils.makePresetEpWeights(
	'P2 - Frost',
	Stats.fromMap(
		{
			[Stat.StatMana]: 0,
			[Stat.StatIntellect]: 0.22,
			[Stat.StatSpirit]: 0.11,
			[Stat.StatSpellDamage]: 1,
			[Stat.StatFrostDamage]: 1,
			[Stat.StatFireDamage]: 0,
			[Stat.StatSpellHitRating]: 1.4,
			[Stat.StatSpellCritRating]: 0.73,
			[Stat.StatSpellHasteRating]: 1.22,
			[Stat.StatSpellPenetration]: 0,
			[Stat.StatMP5]: 0.02,
		},
		{
			[PseudoStat.PseudoStatSchoolHitPercentFrost]: 1.4,
		},
	),
);

export const Talents = {
	name: 'Blank',
	data: SavedTalents.create({
		talentsString: '',
	}),
};

export const DefaultOptions = FrostMageOptions.create({
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

export const P1_PRESET_BUILD_FROST = PresetUtils.makePresetBuild('P1', {
	group: 'Frost',
	phase: Phase.Phase1,
	gear: P1_BIS_FROST,
	talents: FROST_TALENTS,
	epWeights: P1_EP_PRESET,
	rotationType: APLRotation_Type.TypeAPL,
	rotation: ROTATION_PRESET_FROST,
	settings: P1_PLAYER_SETTINGS,
});

export const P2_PRESET_BUILD_FROST = PresetUtils.makePresetBuild('P2', {
	group: 'Frost',
	phase: Phase.Phase2,
	gear: P2_BIS_FROST,
	talents: FROST_TALENTS,
	epWeights: P2_EP_PRESET,
	rotationType: APLRotation_Type.TypeAPL,
	rotation: ROTATION_PRESET_FROST,
	settings: P2_PLAYER_SETTINGS,
});
