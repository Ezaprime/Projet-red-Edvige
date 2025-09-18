package main

import (
	"bufio"
	"fmt"
	"os"
	"piscine"
	"strings"
	"unicode/utf8"
)

func infoScreen(c *piscine.Character) {
	for {
		fmt.Println("\n=== Informations du personnage ===")

		piscine.DisplayInfo(*c)

		fmt.Println("0) Retour")
		var ch int
		fmt.Print("> ")
		fmt.Scanln(&ch)
		if ch == 0 {
			return
		}
	}
}


func askPlayerName() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("=== Choix du nom du héros ===")
		fmt.Print("Entre ton nom (2–16 caractères) : ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		name = strings.Join(strings.Fields(name), " ")
		if n := utf8.RuneCountInString(name); n >= 2 && n <= 16 {
			fmt.Println("Nom choisi :", name)
			return name
		}
		fmt.Println("Nom invalide. Réessaie.")
	}
}

func main() {
	piscine.AnimateNeonBanner()
	name := askPlayerName()

	c := piscine.InitCharacter(
		name,
		"Lacoste TN",
		1, 100, 40,
		[]string{piscine.ItemStimpak, piscine.ItemToxVial},
	)

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
			infoScreen(&c) 
		case 2:
			piscine.AccessInventory(&c)
		case 3:
			piscine.Merchant(&c)
		case 4:
			piscine.Forge(&c)
		case 5:
			piscine.TrainingFight(&c)
		case 6:
			piscine.StoryMode(&c)
		case 0:
			fmt.Println("Au revoir.")
			os.Exit(0)
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
