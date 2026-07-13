import * as OtherInputs from '../../core/components/inputs/other_inputs';
import { IndividualSimUI, registerSpecConfig } from '../../core/individual_sim_ui';
import { Player } from '../../core/player';
import { PlayerClasses } from '../../core/player_classes';
import { Mage } from '../../core/player_classes/mage';
import { APLListItem, APLRotation, APLRotation_Type } from '../../core/proto/apl';
import { Cooldowns, Faction, ItemSlot, PseudoStat, Race, Spec, Stat } from '../../core/proto/common';
import { DEFAULT_CASTER_GEM_STATS, Stats, UnitStat } from '../../core/proto_utils/stats';
import { DefaultDebuffs, DefaultRaidBuffs, DefaultPartyBuffs, DefaultIndividualBuffs, DefaultConsumables } from './presets';
import { SpecRotation } from '../../core/proto_utils/utils';
import * as AplUtils from '../../core/proto_utils/apl_utils';
import * as Presets from './presets';
import * as MageInputs from './inputs';
import { ReforgeOptimizer } from '../../core/components/suggest_reforges_action';

const SPEC_CONFIG = registerSpecConfig(Spec.SpecFireMage, {
	requiredTalentRows: [],
	cssClass: 'fire-mage-sim-ui',
	cssScheme: PlayerClasses.getCssClass(Mage),
	// List any known bugs / issues here and they'll be shown on the site.
	knownIssues: [],

	// All stats for which EP should be calculated.
	epStats: [
		Stat.StatIntellect,
		Stat.StatSpirit,
		Stat.StatSpellDamage,
		Stat.StatArcaneDamage,
		Stat.StatFrostDamage,
		Stat.StatFireDamage,
		Stat.StatSpellPenetration,
		Stat.StatSpellHitRating,
		Stat.StatSpellCritRating,
		Stat.StatSpellHasteRating,
		Stat.StatMana,
		Stat.StatMP5,
	],
	epPseudoStats: [PseudoStat.PseudoStatSchoolHitPercentArcane, PseudoStat.PseudoStatSchoolHitPercentFire, PseudoStat.PseudoStatSchoolHitPercentFrost],
	// Reference stat against which to calculate EP. I think all classes use either spell power or attack power.
	epReferenceStat: Stat.StatSpellDamage,
	// Which stats to display in the Character Stats section, at the bottom of the left-hand sidebar.
	displayStats: UnitStat.createDisplayStatArray(
		[
			Stat.StatHealth,
			Stat.StatMana,
			Stat.StatStamina,
			Stat.StatIntellect,
			Stat.StatSpirit,
			Stat.StatSpellDamage,
			Stat.StatFrostDamage,
			Stat.StatFireDamage,
			Stat.StatArcaneDamage,
		],
		[
			PseudoStat.PseudoStatSpellHitPercent,
			PseudoStat.PseudoStatSchoolHitPercentArcane,
			PseudoStat.PseudoStatSchoolHitPercentFire,
			PseudoStat.PseudoStatSchoolHitPercentFrost,
			PseudoStat.PseudoStatSpellCritPercent,
			PseudoStat.PseudoStatSpellHastePercent,
		],
	),

	modifyDisplayStats: (player: Player<Spec.SpecFireMage>) => {
		return {
			talents: new Stats().addPseudoStat(PseudoStat.PseudoStatSpellCritPercent, player.getTalents().arcaneInstability),
		};
	},

	gemStats: DEFAULT_CASTER_GEM_STATS,

	consumableStats: [
		Stat.StatIntellect,
		Stat.StatSpirit,
		Stat.StatMP5,
		Stat.StatMana,
		Stat.StatSpellDamage,
		Stat.StatFrostDamage,
		Stat.StatFireDamage,
		Stat.StatArcaneDamage,
		Stat.StatSpellCritRating,
		Stat.StatSpellHitRating,
		Stat.StatSpellHasteRating,
	],

	defaults: {
		// Default equipped gear.
		gear: Presets.P2_BIS_FIRE.gear,
		// Default EP weights for sorting gear in the gear picker.
		epWeights: Presets.P2_EP_PRESET.epWeights,
		statCaps: (() => {
			return new Stats().withPseudoStat(PseudoStat.PseudoStatSchoolHitPercentFire, 16);
		})(),
		// Default consumes settings.
		consumables: DefaultConsumables,
		// Default talents.
		talents: Presets.FIRE_TALENTS.data,
		// Default spec-specific settings.
		specOptions: Presets.DefaultOptions,
		other: Presets.OtherDefaults,
		// Default raid/party buffs settings.
		raidBuffs: DefaultRaidBuffs,

		partyBuffs: DefaultPartyBuffs,
		individualBuffs: DefaultIndividualBuffs,

		rotationType: APLRotation_Type.TypeSimple,
		simpleRotation: Presets.FireMageSimpleRotation,
		debuffs: DefaultDebuffs,
	},

	// IconInputs to include in the 'Player' section on the settings tab.
	playerIconInputs: [MageInputs.MageArmorInputs()],
	rotationInputs: MageInputs.FireMageRotationConfig,
	// Buff and Debuff inputs to include/exclude, overriding the EP-based defaults.
	includeBuffDebuffInputs: [Stat.StatMP5],
	excludeBuffDebuffInputs: [],
	// Inputs to include in the 'Other' section on the settings tab.
	otherInputs: {
		inputs: [OtherInputs.InputDelay, OtherInputs.DistanceFromTarget, OtherInputs.TankAssignment],
	},
	itemSwapSlots: [ItemSlot.ItemSlotMainHand, ItemSlot.ItemSlotOffHand, ItemSlot.ItemSlotTrinket1, ItemSlot.ItemSlotTrinket2],
	encounterPicker: {
		// Whether to include 'Execute Duration (%)' in the 'Encounter' section of the settings tab.
		showExecuteProportion: true,
	},

	presets: {
		epWeights: [Presets.P1_EP_PRESET, Presets.P2_EP_PRESET],
		// Preset rotations that the user can quickly select.
		rotations: [Presets.APL_FIRE_SIMPLE],
		// Preset talents that the user can quickly select.
		talents: [Presets.FIRE_TALENTS],
		// Preset gear configurations that the user can quickly select.
		gear: [Presets.PREBIS_FIRE, Presets.P1_BIS_FIRE, Presets.P2_BIS_FIRE],

		builds: [Presets.P1_PRESET_BUILD_FIRE, Presets.P2_PRESET_BUILD_FIRE],
	},

	autoRotation: (_player: Player<Spec.SpecFireMage>): APLRotation => {
		return Presets.ROTATION_PRESET_FIRE.rotation.rotation!;
	},

	simpleRotation: (_player: Player<Spec.SpecFireMage>, simple: SpecRotation<Spec.SpecFireMage>, cooldowns: Cooldowns): APLRotation => {
		const actions = AplUtils.simpleCooldownActions(cooldowns);
		const rotation = APLRotation.clone(Presets.ROTATION_PRESET_FIRE.rotation.rotation!);

		const priorityList = simple.weaveFireBlast
			? rotation.priorityList
			: rotation.priorityList.filter(
					item =>
						item.action?.action.oneofKind !== 'castSpell' ||
						item.action.action.castSpell.spellId?.rawId.oneofKind !== 'spellId' ||
						item.action.action.castSpell.spellId.rawId.spellId !== 27079,
				);

		return APLRotation.create({
			prepullActions: rotation.prepullActions,
			priorityList: [
				...actions.map(action =>
					APLListItem.create({
						action: action,
					}),
				),
				...priorityList,
			],
			groups: rotation.groups,
			valueVariables: rotation.valueVariables,
		});
	},

	raidSimPresets: [
		{
			spec: Spec.SpecFireMage,
			talents: Presets.Talents.data,
			specOptions: Presets.DefaultOptions,
			consumables: Presets.DefaultConsumables,
			otherDefaults: Presets.OtherDefaults,
			defaultFactionRaces: {
				[Faction.Unknown]: Race.RaceUnknown,
				[Faction.Alliance]: Race.RaceGnome,
				[Faction.Horde]: Race.RaceTroll,
			},
			defaultGear: {
				[Faction.Unknown]: {},
				[Faction.Alliance]: {
					1: Presets.BLANK_GEARSET.gear,
				},
				[Faction.Horde]: {
					1: Presets.BLANK_GEARSET.gear,
				},
			},
		},
	],
});

export class FireMageSimUI extends IndividualSimUI<Spec.SpecFireMage> {
	constructor(parentElem: HTMLElement, player: Player<Spec.SpecFireMage>) {
		super(parentElem, player, SPEC_CONFIG);

		this.reforger = new ReforgeOptimizer(this);
	}
}
