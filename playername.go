package piscine

import (
	"fmt"
)

func AskPlayerName() string {
	var name string

	for {
		fmt.Println("=== Choix du nom du héros ===")
		fmt.Print("Entre ton nom (2–16 caractères) : ")
		fmt.Scanln(&name) 
		
		if len(name) >= 2 && len(name) <= 16 {
			fmt.Println("Nom choisi :", name)
			return name
		} else {
			fmt.Println("Nom invalide. Réessaie.")
		}
	}
}
