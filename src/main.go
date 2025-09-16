package main

import "fmt"

func main() {
	c := initCharacter("Kerem", "Elfe")
	fmt.Println("Personnage créé :", c.Name, "PV :", c.Hp, "/", c.HpMax)
}
