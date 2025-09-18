package piscine

import "fmt"

func Forge(c *Character) {

	for {

		fmt.Println("=== Forgeron : Vuldar ===")

		fmt.Println("Crédits :", c.Money)

		fmt.Println("Recettes disponibles :")

		fmt.Println("1) Casque Nomade   (1 Plume-Corbeau, 1 Cuir-Sanglier) — coût 5")

		fmt.Println("2) Kevlar Nomade   (2 Fibre-Loup, 1 Peau-Troll)      — coût 5")

		fmt.Println("3) Bottes Nomades  (1 Fibre-Loup, 1 Cuir-Sanglier)   — coût 5")

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

		fmt.Println("Sacoche pleine : impossible d'ajouter l'équipement fabriqué.")

		return

	}

	if !takeItems(c, recipe) {

		fmt.Println("Erreur lors du retrait des ressources.")

		return

	}

	c.Money -= ForgeCost

	c.Inventory = append(c.Inventory, equipName)

	fmt.Println("Fabriqué :", equipName, "(ajouté à la sacoche)")

}
