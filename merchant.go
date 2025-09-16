package main

import "fmt"

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
		fmt.Println("9) Augmentation de sacoche (+10 cap, max 3) - 30 crédits")
		fmt.Println("10) Pistolet Neo-9 - 40 crédits")
		fmt.Println("11) SMG Kiroshi - 55 crédits")
		fmt.Println("12) Katana Monofil - 50 crédits")
		fmt.Println("13) Puff (one-shot) - 200 crédits [Niveau 3+]")
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
		case 10:
			buyItem(c, WeaponPistol)
		case 11:
			buyItem(c, WeaponSMG)
		case 12:
			buyItem(c, WeaponKatana)
		case 13:
			buyItem(c, WeaponPuff)
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func buyItem(c *Character, item string) {
	if item == ItemUpgradeInventory && c.InventoryUpgrades >= 3 {
		fmt.Println("Déjà 3 augmentations de sacoche.")
		return
	}
	if item == WeaponPuff && c.Level < 3 {
		fmt.Println("La Puff est réservée aux niveaux 3+. Gagne de l'XP puis reviens.")
		return
	}
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
		fmt.Println("Sacoche pleine. Revends/consomme avant d'acheter.")
		return
	}
	c.Money -= cost
	c.Inventory = append(c.Inventory, item)
	fmt.Println("Achat effectué :", item)
}
