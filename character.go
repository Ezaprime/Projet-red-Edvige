package piscine

import (
	"fmt"

	"strings"

	"time"
)

func InitCharacter(name string, class string, level int, hpMax int, hp int, inv []string) Character {

	return Character{

		Name: name, 

		Class: class,

		Level: level,

		BaseHpMax: hpMax,

		HpMax: hpMax,

		Hp: hp,

		AttackBase: 5,

		Inventory: inv,

		InventoryCap: 10,

		InventoryUpgrades: 0,

		Money: 100,

		Equip: Equipment{},

		XP: 0,

		XPMax: 20,

		ManaMax: 10,

		Mana: 10,

		Skills: []string{SpellPunch},

		StoryProgress: 1,
	}

}

func DisplayInfo(c Character) {

	fmt.Println("=== Informations du personnage ===")

	fmt.Println("Nom :", c.Name)

	fmt.Println("Classe :", c.Class)

	fmt.Println("Niveau :", c.Level)

	fmt.Printf("PV : %d / %d (base %d)\n", c.Hp, c.HpMax, c.BaseHpMax)

	fmt.Println("Attaque de base :", c.AttackBase)

	fmt.Printf("XP : %d / %d\n", c.XP, c.XPMax)

	fmt.Printf("Mana : %d / %d\n", c.Mana, c.ManaMax)

	if len(c.Skills) == 0 {

		fmt.Println("Sorts connus : (aucun)")

	} else {

		fmt.Println("Sorts connus :", strings.Join(c.Skills, ", "))

	}

	fmt.Println("Équipement : Tête[", c.Equip.Head, "] Torse[", c.Equip.Torso, "] Pieds[", c.Equip.Feet, "] Arme[", c.Equip.Weapon, "]")

	fmt.Printf("Sacoche : %d / %d → %v\n", len(c.Inventory), c.InventoryCap, c.Inventory)

	fmt.Println("Crédits :", c.Money)

	if c.StoryProgress >= 1 && c.StoryProgress <= len(chapters) {

		fmt.Printf("Progression histoire : Chapitre %d/%d — %s\n", c.StoryProgress, len(chapters), chapters[c.StoryProgress-1].Title)

	} else if c.StoryProgress > len(chapters) {

		fmt.Println("Progression histoire : Terminée")

	}

	fmt.Println("==================================")

}

func hasSpace(c *Character) bool { return len(c.Inventory) < c.InventoryCap }

func addInventory(c *Character, item string) bool {

	if !hasSpace(c) {

		fmt.Println("Sacoche pleine (capacité atteinte).")

		return false

	}

	c.Inventory = append(c.Inventory, item)

	fmt.Println("Ajouté à la sacoche :", item)

	return true

}

func removeInventoryAt(c *Character, index int) {

	if index < 0 || index >= len(c.Inventory) {

		return

	}

	c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)

}

func clampHP(c *Character) {

	if c.Hp > c.HpMax {

		c.Hp = c.HpMax

	}

	if c.Hp < 0 {

		c.Hp = 0

	}

}

func isDead(c *Character) bool {

	if c.Hp <= 0 {

		fmt.Println(c.Name, "tombe à 0 PV... Résurrection à 50% des PV max.")

		c.Hp = c.HpMax / 2

		return true

	}

	return false

}

func countItem(c *Character, name string) int {

	cnt := 0

	for _, it := range c.Inventory {

		if it == name {

			cnt++

		}

	}

	return cnt

}

func takeItems(c *Character, need map[string]int) bool {

	for k, v := range need {

		if countItem(c, k) < v {

			return false

		}

	}

	for k, v := range need {

		for i := 0; i < v; i++ {

			removeFirst(c, k)

		}

	}

	return true

}

func removeFirst(c *Character, name string) {

	for i, it := range c.Inventory {

		if it == name {

			removeInventoryAt(c, i)

			return

		}

	}

}

func useItemByIndex(c *Character, index int) {

	if index < 0 || index >= len(c.Inventory) {

		fmt.Println("Indice invalide.")

		return

	}

	item := c.Inventory[index]

	switch item {

	case ItemStimpak:

		removeInventoryAt(c, index)

		c.Hp += 50

		clampHP(c)

		fmt.Printf("Vous utilisez un Stimpak. PV : %d / %d\n", c.Hp, c.HpMax)

	case ItemToxVial:

		removeInventoryAt(c, index)

		poisonPot(c)

	case ItemManaBattery:

		removeInventoryAt(c, index)

		c.Mana += 10

		if c.Mana > c.ManaMax {

			c.Mana = c.ManaMax

		}

		fmt.Printf("Vous utilisez une Batterie d'appoint. Mana : %d / %d\n", c.Mana, c.ManaMax)

	case ItemFireballChip:

		removeInventoryAt(c, index)

		spellBook(c)

	case ItemUpgradeInventory:

		removeInventoryAt(c, index)

		if c.InventoryUpgrades >= 3 {

			fmt.Println("Capacité déjà augmentée 3 fois (maximum atteint).")

			return

		}

		c.InventoryCap += 10

		c.InventoryUpgrades++

		fmt.Printf("Capacité de sacoche augmentée : %d (utilisations %d/3)\n", c.InventoryCap, c.InventoryUpgrades)

	case EquipCasqueNomade, EquipKevlarNomade, EquipBottesNomades, WeaponPistol, WeaponSMG, WeaponKatana, WeaponPuff:

		removeInventoryAt(c, index)

		equipItem(c, item)

	default:

		fmt.Println("Cet objet n'est pas utilisable pour le moment.")

	}

}

func spellBook(c *Character) {

	for _, s := range c.Skills {

		if s == SpellFireball {

			fmt.Println("Boule de feu est déjà apprise.")

			return

		}

	}

	c.Skills = append(c.Skills, SpellFireball)

	fmt.Println("Nouveau sort appris :", SpellFireball)

}

func poisonPot(c *Character) {

	fmt.Println("Vous utilisez Tox-Vial : 10 dégâts par seconde pendant 3 secondes.")

	for i := 1; i <= 3; i++ {

		c.Hp -= 10

		clampHP(c)

		fmt.Printf("Seconde %d → PV : %d / %d\n", i, c.Hp, c.HpMax)

		time.Sleep(1 * time.Second)

	}

	if isDead(c) {

		fmt.Printf("Après résurrection → PV : %d / %d\n", c.Hp, c.HpMax)

	}

}

func equipItem(c *Character, item string) {
	var old string
	switch item {
	case EquipCasqueNomade:
		old = c.Equip.Head
		c.Equip.Head = EquipCasqueNomade
	case EquipKevlarNomade:
		old = c.Equip.Torso
		c.Equip.Torso = EquipKevlarNomade
	case EquipBottesNomades:
		old = c.Equip.Feet
		c.Equip.Feet = EquipBottesNomades
	case WeaponPistol, WeaponSMG, WeaponKatana, WeaponPuff:
		old = c.Equip.Weapon
		c.Equip.Weapon = item
	default:
		return
	}
	if old != "" {
		if !addInventory(c, old) {
			fmt.Println("Sacoche pleine, l'ancien équipement est perdu :", old)
		}
	}
	applyEquipBonuses(c)
	fmt.Println("Équipement mis à jour. PV max recalculés :", c.HpMax)
	if c.Hp > c.HpMax {
		c.Hp = c.HpMax
	}
}

func applyEquipBonuses(c *Character) {

	bonus := 0

	if c.Equip.Head == EquipCasqueNomade {

		bonus += 10

	}

	if c.Equip.Torso == EquipKevlarNomade {

		bonus += 25

	}

	if c.Equip.Feet == EquipBottesNomades {

		bonus += 15

	}

	c.HpMax = c.BaseHpMax + bonus

}

func AccessInventory(c *Character) {

	for {

		fmt.Printf("=== Sacoche (%d / %d) ===\n", len(c.Inventory), c.InventoryCap)

		if len(c.Inventory) == 0 {

			fmt.Println("(Sacoche vide)")

		} else {

			for i, item := range c.Inventory {

				fmt.Printf("%d) %s\n", i+1, item)

			}

		}

		fmt.Println("0) Retour")

		var choice int

		fmt.Print("> ")

		fmt.Scanln(&choice)

		if choice == 0 {

			return

		}

		index := choice - 1

		if index < 0 || index >= len(c.Inventory) {

			fmt.Println("Choix invalide.")

			continue

		}

		useItemByIndex(c, index)

	}

}
