package piscine

import (
	"fmt"
	"time"
)

func InitGoblin() Monster { return Monster{Name: "Sparring holo", HpMax: 40, Hp: 40, Attack: 5} }
func initMunicipale() Monster {
	return Monster{Name: "Police municipale", HpMax: 35, Hp: 35, Attack: 6}
}
func initNationale() Monster   { return Monster{Name: "Police nationale", HpMax: 55, Hp: 55, Attack: 7} }
func initGendarmerie() Monster { return Monster{Name: "Gendarmerie", HpMax: 70, Hp: 70, Attack: 8} }
func initMotards() Monster     { return Monster{Name: "Motards", HpMax: 85, Hp: 85, Attack: 9} }
func initBAC() Monster         { return Monster{Name: "BAC", HpMax: 100, Hp: 100, Attack: 10} }
func initBossFBI() Monster     { return Monster{Name: "FBI", HpMax: 220, Hp: 220, Attack: 12} }

func EnemyPattern(g *Monster, c *Character, turn int) {
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

func charTurn(c *Character, g *Monster, bossFight bool) {
	for {
		fmt.Println("— Tour du joueur —")
		fmt.Println("1) Attaquer (attaque basique)")
		fmt.Println("2) Sacoche (utiliser un objet)")
		fmt.Println("3) Sorts")
		fmt.Println("4) Attaquer avec l'arme")
		fmt.Println("0) Passer / Retour")
		var choice int
		fmt.Print("> ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			if bossFight {
				fmt.Println("Votre attaque de base est inefficace contre", g.Name, ".")
				return
			}
			dmg := c.AttackBase
			g.Hp -= dmg
			if g.Hp < 0 {
				g.Hp = 0
			}
			fmt.Printf("%s utilise Attaque basique → %d dégâts → PV ennemi: %d/%d\n", c.Name, dmg, g.Hp, g.HpMax)
			return

		case 2:
			if len(c.Inventory) == 0 {
				fmt.Println("(Sacoche vide)")
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
			if bossFight {
				fmt.Println(g.Name, "annule votre sort :", sel, ".")
				return
			}
			c.Mana -= cost
			dmg := spellDamage[sel]
			g.Hp -= dmg
			if g.Hp < 0 {
				g.Hp = 0
			}
			fmt.Printf("%s lance %s → %d dégâts → PV ennemi: %d/%d (Mana %d/%d)\n",
				c.Name, sel, dmg, g.Hp, g.HpMax, c.Mana, c.ManaMax)
			return

		case 4:
			if c.Equip.Weapon == "" {
				fmt.Println("Aucune arme équipée.")
				continue
			}
			if bossFight {
				if c.Equip.Weapon != WeaponPuff {
					fmt.Println("Cette arme ne passe pas. Il faut la Puff.")
					return
				}
				fmt.Printf("%s dégaine la Puff... %s s'effondre instantanément !\n", c.Name, g.Name)
				g.Hp = 0
				fmt.Printf("PV ennemi: %d/%d\n", g.Hp, g.HpMax)
				return
			}
			dmg, ok := weaponDamage[c.Equip.Weapon]
			if !ok {
				fmt.Println("Arme inconnue.")
				continue
			}
			g.Hp -= dmg
			if g.Hp < 0 {
				g.Hp = 0
			}
			action := "tire avec son arme"
			if c.Equip.Weapon == WeaponKatana {
				action = "frappe avec son katana"
			}
			fmt.Printf("%s %s (%s) → %d dégâts → PV ennemi: %d/%d\n", c.Name, action, c.Equip.Weapon, dmg, g.Hp, g.HpMax)
			return

		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func rollInitiative() int { return int(time.Now().UnixNano()%20) + 1 }

func grantXP(c *Character, amount int) {
	c.XP += amount
	fmt.Printf("Vous gagnez %d XP (total %d/%d).\n", amount, c.XP, c.XPMax)
	for c.XP >= c.XPMax {
		c.XP -= c.XPMax
		levelUp(c)
	}
	fmt.Printf("XP actuel: %d/%d\n", c.XP, c.XPMax)
}

func knows(c *Character, spell string) bool {
	for _, s := range c.Skills {
		if s == spell {
			return true
		}
	}
	return false
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
	if c.Level >= 2 && !knows(c, SpellFireball) {
		c.Skills = append(c.Skills, SpellFireball)
		fmt.Println("Nouvelle maîtrise : Boule de feu (débloquée par niveau 2).")
	}
}

func TrainingFight(c *Character) {
	idx := c.StoryProgress - 1
	if idx < 0 {
		idx = 0
	}
	if idx >= len(chapters) {
		idx = len(chapters) - 1
	}
	baseXP := chapters[idx].XpReward
	baseCr := chapters[idx].CreditReward
	xpReward := baseXP / 4
	if xpReward < 1 {
		xpReward = 1
	}
	creditReward := baseCr / 4
	if creditReward < 1 {
		creditReward = 1
	}

	g := InitGoblin()
	c.Initiative = rollInitiative()
	g.Initiative = rollInitiative()
	playerFirst := c.Initiative >= g.Initiative
	turn := 1
	fmt.Println("=== Entraînement —", g.Name, "===")
	for {
		fmt.Println("--- Tour", turn, "---")
		if playerFirst {
			charTurn(c, &g, false)
			if g.Hp <= 0 {
				fmt.Println(g.Name, "est vaincu !")
				grantXP(c, xpReward)
				c.Money += creditReward
				fmt.Printf("+%d crédits (total %d)\n", creditReward, c.Money)
				fmt.Println("Retour au menu.")
				return
			}
			EnemyPattern(&g, c, turn)
			if c.Hp <= 0 {
				fmt.Println("Vous avez été mis K.O. Retour au menu.")
				return
			}
		} else {
			EnemyPattern(&g, c, turn)
			if c.Hp <= 0 {
				fmt.Println("Vous avez été mis K.O. Retour au menu.")
				return
			}
			charTurn(c, &g, false)
			if g.Hp <= 0 {
				fmt.Println(g.Name, "est vaincu !")
				grantXP(c, xpReward)
				c.Money += creditReward
				fmt.Printf("+%d crédits (total %d)\n", creditReward, c.Money)
				fmt.Println("Retour au menu.")
				return
			}
		}
		turn++
	}
}

func fightGeneric(c *Character, m Monster, xpReward int, creditReward int) bool {
	g := m
	c.Initiative = rollInitiative()
	g.Initiative = rollInitiative()
	playerFirst := c.Initiative >= g.Initiative
	turn := 1
	fmt.Println("=== Combat —", g.Name, "===")
	for {
		fmt.Println("--- Tour", turn, "---")
		if playerFirst {
			charTurn(c, &g, false)
			if g.Hp <= 0 {
				fmt.Println(g.Name, "est vaincu !")
				grantXP(c, xpReward)
				c.Money += creditReward
				fmt.Printf("+%d crédits (total %d)\n", creditReward, c.Money)
				return true
			}
			EnemyPattern(&g, c, turn)
			if c.Hp <= 0 {
				fmt.Println("Vous êtes K.O.")
				return false
			}
		} else {
			EnemyPattern(&g, c, turn)
			if c.Hp <= 0 {
				fmt.Println("Vous êtes K.O.")
				return false
			}
			charTurn(c, &g, false)
			if g.Hp <= 0 {
				fmt.Println(g.Name, "est vaincu !")
				grantXP(c, xpReward)
				c.Money += creditReward
				fmt.Printf("+%d crédits (total %d)\n", creditReward, c.Money)
				return true
			}
		}
		turn++
	}
}

func fightBoss(c *Character, xpReward int, creditReward int) bool {
	g := initBossFBI()
	c.Initiative = rollInitiative()
	g.Initiative = rollInitiative()
	playerFirst := c.Initiative >= g.Initiative
	turn := 1
	fmt.Println("=== BOSS FINAL —", g.Name, "===")
	for {
		fmt.Println("--- Tour", turn, "---")
		if playerFirst {
			charTurn(c, &g, true)
			if g.Hp <= 0 {
				fmt.Println(g.Name, "est pulvérisé !")
				grantXP(c, xpReward)
				c.Money += creditReward
				fmt.Printf("+%d crédits (total %d)\n", creditReward, c.Money)
				return true
			}
			EnemyPattern(&g, c, turn)
			if c.Hp <= 0 {
				fmt.Println("Vous êtes K.O.")
				return false
			}
		} else {
			EnemyPattern(&g, c, turn)
			if c.Hp <= 0 {
				fmt.Println("Vous êtes K.O.")
				return false
			}
			charTurn(c, &g, true)
			if g.Hp <= 0 {
				fmt.Println(g.Name, "est pulvérisé !")
				grantXP(c, xpReward)
				c.Money += creditReward
				fmt.Printf("+%d crédits (total %d)\n", creditReward, c.Money)
				return true
			}
		}
		turn++
	}
}

func StoryMode(c *Character) {
	if c.StoryProgress > len(chapters) {
		fmt.Println("Histoire déjà terminée. Retour au menu.")
		return
	}
	ch := chapters[c.StoryProgress-1]
	fmt.Println("=== MODE HISTOIRE : CYBERYNOV — Rat Lyonnais ===")
	fmt.Println("Chapitre", c.StoryProgress, "—", ch.Title)

	if ch.IsBoss {
		if c.Level < 3 {
			fmt.Println("Niveau insuffisant (3+ requis) pour affronter le boss. Entraîne-toi.")
			return
		}
		if c.Equip.Weapon != WeaponPuff && countItem(c, WeaponPuff) == 0 {
			fmt.Printf("Il te faut la Puff (coût %d). Achète-la chez Mox et équipe-la.\n", prices[WeaponPuff])
			return
		}
		if c.Equip.Weapon != WeaponPuff {
			fmt.Println("Tu possèdes la Puff mais elle n'est pas équipée. Équipe-la dans la sacoche.")
			return
		}
		if fightBoss(c, ch.XpReward, ch.CreditReward) {
			c.StoryProgress++
			if c.StoryProgress > len(chapters) {
				fmt.Println("\n=== FIN — Tu as dominé la scène. Gloire à Edvige ! ===")
			}
			fmt.Println("Retour au menu.")
			return
		}
		fmt.Println("Échec du boss. Retour au menu.")
		return
	}

	if fightGeneric(c, ch.Factory(), ch.XpReward, ch.CreditReward) {
		fmt.Println("Chapitre réussi !")
		c.StoryProgress++
		fmt.Println("Retour au menu.")
	} else {
		fmt.Println("Échec du chapitre. Retour au menu.")
	}
}
