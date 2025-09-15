package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Println("\n=== Cyberynov ===")
		fmt.Println("0) Quitter")

		var choice int
		fmt.Print("> ")
		fmt.Scanln(&choice)

		switch choice {
		case 0:
			fmt.Println("Au revoir.")
			os.Exit(0)
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
