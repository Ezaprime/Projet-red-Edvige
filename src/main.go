package main

import "fmt"

func main() {
	c := initCharacter("Kerem", "Elfe")
	addInventory(&c, "Stimpak")
	addInventory(&c, "Tox-Vial")
	fmt.Println("Inventaire :", c.Inventory)
}
