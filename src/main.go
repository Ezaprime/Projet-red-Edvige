package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	ItemStimpak          = "Stimpak"
	ItemToxVial          = "Tox-Vial"
	ItemManaBattery      = "Batterie d'appoint"
	ItemFireballChip     = "Chip de sort : Fireball"
	ItemFibreLoup        = "Fibre-Loup"
	ItemPeauTroll        = "Peau-Troll"
	ItemCuirSanglier     = "Cuir-Sanglier"
	ItemPlumeCorbeau     = "Plume-Corbeau"
	ItemUpgradeInventory = "Augmentation d'inventaire"
	EquipCasqueNomade    = "Casque Nomade"
	EquipKevlarNomade    = "Kevlar Nomade"
	EquipBottesNomades   = "Bottes Nomades"
)

var prices = map[string]int{
	ItemStimpak:          3,
	ItemToxVial:          6,
	ItemManaBattery:      5,
	ItemFireballChip:     25,
	ItemFibreLoup:        4,
	ItemPeauTroll:        7,
	ItemCuirSanglier:     3,
	ItemPlumeCorbeau:     1,
	ItemUpgradeInventory: 30,
}

const ForgeCost = 5

var recipeCasque = map[string]int{ItemPlumeCorbeau: 1, ItemCuirSanglier: 1}
var recipeKevlar = map[string]int{ItemFibreLoup: 2, ItemPeauTroll: 1}
var recipeBottes = map[string]int{ItemFibreLoup: 1, ItemCuirSanglier: 1}

const (
	SpellPunch    = "Coup de poing"
	SpellFireball = "Boule de feu"
)

var manaCost = map[string]int{
	SpellPunch:    2,
	SpellFireball: 5,
}

var spellDamage = map[string]int{
	SpellPunch:    8,
	SpellFireball: 18,
}

type Equipment struct {
	Head  string
	Torso string
	Feet  string
}

type Character struct {
	Name              string
	Class             string
	Level             int
	BaseHpMax         int
	HpMax             int
	Hp                int
	AttackBase        int
	Inventory         []string
	InventoryCap      int
	InventoryUpgrades int
	Money             int
	Equip             Equipment
	Initiative        int
	XP                int
	XPMax             int
	ManaMax           int
	Mana              int
	Skills            []string
}

type Monster struct {
	Name       string
	HpMax      int
	Hp         int
	Attack     int
	Initiative int
}

func initCharacter(name string, class string, level int, hpMax int, hp int, inv []string) Character {
	return Character{
		Name:              name,
		Class:             class,
		Level:             level,
		BaseHpMax:         hpMax,
		HpMax:             hpMax,
		Hp:                hp,
		AttackBase:        5,
		Inventory:         inv,
		InventoryCap:      10,
		InventoryUpgrades: 0,
		Money:             100,
		Equip:             Equipment{},
		XP:                0,
		XPMax:             20,
		ManaMax:           10,
		Mana:              10,
		Skills:            []string{SpellPunch},
	}
}

func displayInfo(c Character) {
	fmt.Println("=== Informations du personnage ===")
	fmt.Println("Nom :", c.Name)
	fmt.Println("Classe :", c.Class)
	fmt.Println("Niveau :", c.Level)
	fmt.Printf("PV : %d / %d (base %d)\n", c.Hp, c.HpMax, c.BaseHpMax)
	fmt.Println("Attaque de base :", c.AttackBase)
	fmt.Printf("XP : %d / %d\n", c.XP, c.XPMax)
	fmt.Printf("Mana : %d / %d\n", c.Mana, c.ManaMax)
	fmt.Println("Sorts connus :", strings.Join(c.Skills, ", "))
	fmt.Println("Équipement : Tête[", c.Equip.Head, "] Torse[", c.Equip.Torso, "] Pieds[", c.Equip.Feet, "]")
	fmt.Printf("Inventaire : %d / %d → %v\n", len(c.Inventory), c.InventoryCap, c.Inventory)
	fmt.Println("Crédits :", c.Money)
	fmt.Println("==================================")
}

func hasSpace(c *Character) bool {
	return len(c.Inventory) < c.InventoryCap
}

func addInventory(c *Character, item string) bool {
	if !hasSpace(c) {
		fmt.Println("Inventaire plein (capacité atteinte).")
		return false
	}
	c.Inventory = append(c.Inventory, item)
	fmt.Println("Ajouté à l'inventaire :", item)
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
		fmt.Printf("Capacité d'inventaire augmentée : %d (utilisations %d/3)\n", c.InventoryCap, c.InventoryUpgrades)
	case EquipCasqueNomade, EquipKevlarNomade, EquipBottesNomades:
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
	switch item {
	case EquipCasqueNomade:
		old := c.Equip.Head
		c.Equip.Head = EquipCasqueNomade
		if old != "" {
			if !addInventory(c, old) {
				fmt.Println("Inventaire plein, l'ancien équipement est perdu :", old)
			}
		}
	case EquipKevlarNomade:
		old := c.Equip.Torso
		c.Equip.Torso = EquipKevlarNomade
		if old != "" {
			if !addInventory(c, old) {
				fmt.Println("Inventaire plein, l'ancien équipement est perdu :", old)
			}
		}
	case EquipBottesNomades:
		old := c.Equip.Feet
		c.Equip.Feet = EquipBottesNomades
		if old != "" {
			if !addInventory(c, old) {
				fmt.Println("Inventaire plein, l'ancien équipement est perdu :", old)
			}
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

func accessInventory(c *Character) {
	for {
		fmt.Printf("=== Inventaire (%d / %d) ===\n", len(c.Inventory), c.InventoryCap)
		if len(c.Inventory) == 0 {
			fmt.Println("(Inventaire vide)")
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

func merchant(c *Character) {
	for {
		fmt.Println("=== Marchand : Mox ===")
		fmt.Println("Crédits disponibles :", c.Money)
		fmt.Println("1) Stimpak - 3 crédits")
		fmt.Println("2) Tox-Vial - 6 crédits")
		fmt.Println("3) Batterie d'appoint - 5 crédits")
		fmt.Println("4) Chip de sort : Fireball - 25 crédits")
		fmt.Println("5) Fibre-Loup - 4 crédits")
		fmt.Println("6) Peau-Troll - 7 crédits")
		fmt.Println("7) Cuir-Sanglier - 3 crédits")
		fmt.Println("8) Plume-Corbeau - 1 crédit")
		fmt.Println("9) Augmentation d'inventaire (+10 cap, max 3) - 30 crédits")
		fmt.Println("0) Retour")

		var choice int
		fmt.Print("> ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			buyItem(c, ItemStimpak)
		case 2:
			buyItem(c, ItemToxVial)
		case 3:
			buyItem(c, ItemManaBattery)
		case 4:
			buyItem(c, ItemFireballChip)
		case 5:
			buyItem(c, ItemFibreLoup)
		case 6:
			buyItem(c, ItemPeauTroll)
		case 7:
			buyItem(c, ItemCuirSanglier)
		case 8:
			buyItem(c, ItemPlumeCorbeau)
		case 9:
			buyItem(c, ItemUpgradeInventory)
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func buyItem(c *Character, item string) {
	cost, ok := prices[item]
	if !ok {
		fmt.Println("Objet inconnu.")
		return
	}
	if c.Money < cost {
		fmt.Println("Pas assez de crédits pour acheter", item)
		return
	}
	if !hasSpace(c) {
		fmt.Println("Inventaire plein. Vendez/consommez des objets avant d'acheter.")
		return
	}
	c.Money -= cost
	c.Inventory = append(c.Inventory, item)
	fmt.Println("Achat effectué :", item)
}

func forge(c *Character) {
	for {
		fmt.Println("=== Forgeron : Vuldar ===")
		fmt.Println("Crédits :", c.Money)
		fmt.Println("Recettes disponibles :")
		fmt.Println("1) Casque Nomade   (recette: 1 Plume-Corbeau, 1 Cuir-Sanglier) — coût 5")
		fmt.Println("2) Kevlar Nomade   (recette: 2 Fibre-Loup, 1 Peau-Troll)      — coût 5")
		fmt.Println("3) Bottes Nomades  (recette: 1 Fibre-Loup, 1 Cuir-Sanglier)   — coût 5")
		fmt.Println("0) Retour")
		var choice int
		fmt.Print("> ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			craftEquip(c, EquipCasqueNomade, recipeCasque)
		case 2:
			craftEquip(c, EquipKevlarNomade, recipeKevlar)
		case 3:
			craftEquip(c, EquipBottesNomades, recipeBottes)
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func canCraft(c *Character, recipe map[string]int) bool {
	for k, v := range recipe {
		if countItem(c, k) < v {
			return false
		}
	}
	return true
}

func craftEquip(c *Character, equipName string, recipe map[string]int) {
	if c.Money < ForgeCost {
		fmt.Println("Crédits insuffisants pour la forge.")
		return
	}
	if !canCraft(c, recipe) {
		fmt.Println("Ressources insuffisantes pour fabriquer", equipName)
		return
	}
	if !hasSpace(c) {
		fmt.Println("Inventaire plein : impossible d'ajouter l'équipement fabriqué.")
		return
	}
	if !takeItems(c, recipe) {
		fmt.Println("Erreur lors du retrait des ressources.")
		return
	}
	c.Money -= ForgeCost
	c.Inventory = append(c.Inventory, equipName)
	fmt.Println("Fabriqué :", equipName, "(ajouté à l'inventaire)")
}

func initGoblin() Monster {
	return Monster{
		Name:   "Gobelin d'entraînement (GOB-LN/Train)",
		HpMax:  40,
		Hp:     40,
		Attack: 5,
	}
}

func goblinPattern(g *Monster, c *Character, turn int) {
	dmg := g.Attack
	if turn%3 == 0 {
		dmg = g.Attack * 2
	}
	c.Hp -= dmg
	clampHP(c)
	fmt.Printf("%s inflige à %s %d de dégâts → PV %s: %d/%d\n", g.Name, c.Name, dmg, c.Name, c.Hp, c.HpMax)
	if c.Hp <= 0 {
		if isDead(c) {
			fmt.Printf("Après résurrection → PV: %d/%d\n", c.Hp, c.HpMax)
		}
	}
}

func charTurn(c *Character, g *Monster) {
	for {
		fmt.Println("— Tour du joueur —")
		fmt.Println("1) Attaquer (attaque basique)")
		fmt.Println("2) Inventaire (utiliser un objet)")
		fmt.Println("3) Sorts")
		fmt.Println("0) Passer / Retour")
		var choice int
		fmt.Print("> ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			dmg := c.AttackBase
			g.Hp -= dmg
			if g.Hp < 0 {
				g.Hp = 0
			}
			fmt.Printf("%s utilise Attaque basique → %d dégâts → PV monstre: %d/%d\n", c.Name, dmg, g.Hp, g.HpMax)
			return
		case 2:
			if len(c.Inventory) == 0 {
				fmt.Println("(Inventaire vide)")
				continue
			}
			for i, it := range c.Inventory {
				fmt.Printf("%d) %s\n", i+1, it)
			}
			fmt.Println("0) Annuler")
			var pick int
			fmt.Print("> ")
			fmt.Scanln(&pick)
			if pick == 0 {
				continue
			}
			idx := pick - 1
			if idx < 0 || idx >= len(c.Inventory) {
				fmt.Println("Choix invalide.")
				continue
			}
			useItemByIndex(c, idx)
			return
		case 3:
			if len(c.Skills) == 0 {
				fmt.Println("Aucun sort connu.")
				continue
			}
			for i, s := range c.Skills {
				fmt.Printf("%d) %s (Coût mana: %d, Dégâts: %d)\n", i+1, s, manaCost[s], spellDamage[s])
			}
			fmt.Println("0) Annuler")
			var sp int
			fmt.Print("> ")
			fmt.Scanln(&sp)
			if sp == 0 {
				continue
			}
			idx := sp - 1
			if idx < 0 || idx >= len(c.Skills) {
				fmt.Println("Choix invalide.")
				continue
			}
			sel := c.Skills[idx]
			cost := manaCost[sel]
			if c.Mana < cost {
				fmt.Println("Mana insuffisant.")
				continue
			}
			c.Mana -= cost
			dmg := spellDamage[sel]
			g.Hp -= dmg
			if g.Hp < 0 {
				g.Hp = 0
			}
			fmt.Printf("%s lance %s → %d dégâts → PV monstre: %d/%d (Mana %d/%d)\n",
				c.Name, sel, dmg, g.Hp, g.HpMax, c.Mana, c.ManaMax)
			return
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func rollInitiative() int {
	n := int(time.Now().UnixNano()%20) + 1
	return n
}

func grantXP(c *Character, amount int) {
	c.XP += amount
	fmt.Printf("Vous gagnez %d XP (total %d/%d).\n", amount, c.XP, c.XPMax)
	for c.XP >= c.XPMax {
		c.XP -= c.XPMax
		levelUp(c)
	}
	fmt.Printf("XP actuel: %d/%d\n", c.XP, c.XPMax)
}

func levelUp(c *Character) {
	c.Level++
	c.AttackBase += 1
	c.BaseHpMax += 10
	applyEquipBonuses(c)
	c.Hp = c.HpMax
	c.ManaMax += 5
	c.Mana = c.ManaMax
	c.XPMax += 10
	fmt.Printf("Niveau %d atteint ! (+10 PV max base, +1 Attaque, +5 Mana max)\n", c.Level)
}

func trainingFight(c *Character) {
	g := initGoblin()
	c.Initiative = rollInitiative()
	g.Initiative = rollInitiative()
	fmt.Printf("Initiatives — %s: %d, %s: %d\n", c.Name, c.Initiative, g.Name, g.Initiative)
	playerFirst := c.Initiative >= g.Initiative
	turn := 1
	fmt.Println("=== Entraînement contre", g.Name, "===")
	for {
		fmt.Println("--- Tour", turn, "---")
		if playerFirst {
			charTurn(c, &g)
			if g.Hp <= 0 {
				fmt.Println(g.Name, "est vaincu !")
				grantXP(c, 20)
				fmt.Println("Combat terminé. Retour au menu.")
				return
			}
			goblinPattern(&g, c, turn)
			if c.Hp <= 0 {
				fmt.Println("Vous avez été mis K.O. Combat terminé. Retour au menu.")
				return
			}
		} else {
			goblinPattern(&g, c, turn)
			if c.Hp <= 0 {
				fmt.Println("Vous avez été mis K.O. Combat terminé. Retour au menu.")
				return
			}
			charTurn(c, &g)
			if g.Hp <= 0 {
				fmt.Println(g.Name, "est vaincu !")
				grantXP(c, 20)
				fmt.Println("Combat terminé. Retour au menu.")
				return
			}
		}
		turn++
	}
}

func main() {
	c := initCharacter("Edvige", "Elfe", 1, 100, 40, []string{
		ItemStimpak, ItemToxVial,
	})

	for {
		fmt.Println("\n=== Menu Principal ===")
		fmt.Println("1) Afficher les infos du personnage")
		fmt.Println("2) Accéder à l'inventaire")
		fmt.Println("3) Marchand")
		fmt.Println("4) Forgeron")
		fmt.Println("5) Entraînement (combat)")
		fmt.Println("0) Quitter")
		var choice int
		fmt.Print("> ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			displayInfo(c)
		case 2:
			accessInventory(&c)
		case 3:
			merchant(&c)
		case 4:
			forge(&c)
		case 5:
			trainingFight(&c)
		case 0:
			fmt.Println("Au revoir.")
			os.Exit(0)
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
