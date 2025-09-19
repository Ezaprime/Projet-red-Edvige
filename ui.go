package piscine

import (
	"fmt"
	"time"
)

func SimpleBanner() {
	// Bannière ASCII de base
	banner := []string{
		" ####   #   #  ####   #####  ####   #   #  #   #   ###   #   #",
		"#         #    #   #  #      #   #   # #   ##  #  #   #  #   #",
		"#       #   #  ####   ####   ####   #   #  # # #  #   #  # # ",
		"#       #   #  #   #  #      #  #   #   #  #  ##  #   #    #  ",
		" ####   #   #  ####   #####  #   #  #   #  #   #   ###     #  ",
		"                 C  Y  B  E  R  Y  N  O  V",
	}

	// Affiche chaque ligne une par une avec une petite pause
	for _, line := range banner {
		fmt.Println(line)
		time.Sleep(200 * time.Millisecond) // pause 0,2 seconde
	}
}

func main() {
	SimpleBanner()
}
