package main

import (
	"fmt"
	"os"
)

func main() {
	c := initCharacter("Edvige", "Elfe", 1, 100, 40, []string{ItemStimpak, ItemToxVial})

	for {
		fmt.Println("\n=== Menu Principal ===")
		fmt.Println("1) Afficher les infos du personnage")
		fmt.Println("2) Ouvrir la sacoche")
		fmt.Println("3) Marchand")
		fmt.Println("4) Forgeron")
		fmt.Println("5) Entraînement (combat)")
		fmt.Println("6) Mode Histoire (chapitre actuel)")
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
		case 6:
			storyMode(&c)
		case 0:
			fmt.Println("Au revoir.")
			os.Exit(0)
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
