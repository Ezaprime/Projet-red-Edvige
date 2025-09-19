package piscine

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func AskPlayerName() string {
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
