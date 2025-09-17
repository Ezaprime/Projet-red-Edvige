package main

const (
	ItemStimpak = "Stimpak"

	ItemToxVial = "Tox-Vial"

	ItemManaBattery = "Batterie d'appoint"

	ItemFireballChip = "Chip de sort : Fireball"

	ItemFibreLoup = "Fibre-Loup"

	ItemPeauTroll = "Peau-Troll"

	ItemCuirSanglier = "Cuir-Sanglier"

	ItemPlumeCorbeau = "Plume-Corbeau"

	ItemUpgradeInventory = "Augmentation d'inventaire"

	EquipCasqueNomade = "Casque Nomade"

	EquipKevlarNomade = "Kevlar Nomade"

	EquipBottesNomades = "Bottes Nomades"

	WeaponPistol = "Pistolet Neo-9"

	WeaponSMG = "SMG Kiroshi"

	WeaponKatana = "Katana Monofil"

	WeaponPuff = "Puff"
)

var prices = map[string]int{

	ItemStimpak: 3,

	ItemToxVial: 6,

	ItemManaBattery: 5,

	ItemFireballChip: 25,

	ItemFibreLoup: 4,

	ItemPeauTroll: 7,

	ItemCuirSanglier: 3,

	ItemPlumeCorbeau: 1,

	ItemUpgradeInventory: 30,

	WeaponPistol: 40,

	WeaponSMG: 55,

	WeaponKatana: 50,

	WeaponPuff: 200,
}

const ForgeCost = 5

var recipeCasque = map[string]int{ItemPlumeCorbeau: 1, ItemCuirSanglier: 1}

var recipeKevlar = map[string]int{ItemFibreLoup: 2, ItemPeauTroll: 1}

var recipeBottes = map[string]int{ItemFibreLoup: 1, ItemCuirSanglier: 1}

const (
	SpellPunch = "Coup de poing"

	SpellFireball = "Boule de feu"
)

var manaCost = map[string]int{

	SpellPunch: 2,

	SpellFireball: 5,
}

var spellDamage = map[string]int{

	SpellPunch: 8,

	SpellFireball: 18,
}

var weaponDamage = map[string]int{

	WeaponPistol: 12,

	WeaponSMG: 9,

	WeaponKatana: 14,

	WeaponPuff: 999999,
}
