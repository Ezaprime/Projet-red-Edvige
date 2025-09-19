package piscine

import (
	"fmt"
	"time"
)

func SimpleBanner() {
	banner := []string{
		" ####   #   #  ####   #####  ####   #   #  #   #   ###   #   #",
		"#         #    #   #  #      #   #   # #   ##  #  #   #  #   #",
		"#       #   #  ####   ####   ####   #   #  # # #  #   #  # # ",
		"#       #   #  #   #  #      #  #   #   #  #  ##  #   #    #  ",
		" ####   #   #  ####   #####  #   #  #   #  #   #   ###     #  ",
		"                 C  Y  B  E  R  Y  N  O  V",
	}

	for _, line := range banner {
		fmt.Println(line)
		time.Sleep(200 * time.Millisecond) 
	}
}

func main() {
	SimpleBanner()
}
