package piscine

import "fmt"

func hasSpace(c *Character) bool {
	return len(c.Inventory) < c.InventoryCap
}
func AddInventory(c *Character, item string) bool {
	if !hasSpace(c) {
		fmt.Println("Sacoche pleine :", item)
		return false
	}
	c.Inventory = append(c.Inventory, item)
	fmt.Println("Ajouté à la sacoche :", item)
	return true
}
func RemoveInventoryAt(c *Character, index int) {
	if index < 0 || index >= len(c.Inventory) {
		return
	}
	c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)
}
