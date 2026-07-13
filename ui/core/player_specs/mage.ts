import { IconSize } from '../player_class';
import { PlayerSpec } from '../player_spec';
import { Class, Spec } from '../proto/common';
import { getSpecSiteUrl } from '../proto_utils/utils';

export class Mage extends PlayerSpec<Spec.SpecMage> {
	static specIndex = 0;
	static specID = Spec.SpecMage as Spec.SpecMage;
	static classID = Class.ClassMage as Class.ClassMage;
	static friendlyName = 'Arcane';
	static simLink = getSpecSiteUrl('mage', 'arcane');

	static isTankSpec = false;
	static isHealingSpec = false;
	static isRangedDpsSpec = true;
	static isMeleeDpsSpec = false;

	static canDualWield = false;

	readonly specIndex = Mage.specIndex;
	readonly specID = Mage.specID;
	readonly classID = Mage.classID;
	readonly friendlyName = Mage.friendlyName;
	readonly simLink = Mage.simLink;

	readonly isTankSpec = Mage.isTankSpec;
	readonly isHealingSpec = Mage.isHealingSpec;
	readonly isRangedDpsSpec = Mage.isRangedDpsSpec;
	readonly isMeleeDpsSpec = Mage.isMeleeDpsSpec;

	readonly canDualWield = Mage.canDualWield;

	static getIcon = (size: IconSize): string => {
		return `https://wow.zamimg.com/images/wow/icons/${size}/spell_holy_magicalsentry.jpg`;
	};

	getIcon = (size: IconSize): string => {
		return Mage.getIcon(size);
	};
}

export class FireMage extends PlayerSpec<Spec.SpecFireMage> {
	static specIndex = 1;
	static specID = Spec.SpecFireMage as Spec.SpecFireMage;
	static classID = Class.ClassMage as Class.ClassMage;
	static friendlyName = 'Fire';
	static simLink = getSpecSiteUrl('mage', 'fire');

	static isTankSpec = false;
	static isHealingSpec = false;
	static isRangedDpsSpec = true;
	static isMeleeDpsSpec = false;

	static canDualWield = false;

	readonly specIndex = FireMage.specIndex;
	readonly specID = FireMage.specID;
	readonly classID = FireMage.classID;
	readonly friendlyName = FireMage.friendlyName;
	readonly simLink = FireMage.simLink;

	readonly isTankSpec = FireMage.isTankSpec;
	readonly isHealingSpec = FireMage.isHealingSpec;
	readonly isRangedDpsSpec = FireMage.isRangedDpsSpec;
	readonly isMeleeDpsSpec = FireMage.isMeleeDpsSpec;

	readonly canDualWield = FireMage.canDualWield;

	static getIcon = (size: IconSize): string => {
		return `https://wow.zamimg.com/images/wow/icons/${size}/spell_fire_firebolt02.jpg`;
	};

	getIcon = (size: IconSize): string => {
		return FireMage.getIcon(size);
	};
}

