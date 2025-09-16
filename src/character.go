package main

type Character struct {
	Name              string
	Class             string
	Level             int
	BaseHpMax         int
	HpMax             int
	Hp                int
	AttackBase        int
	Inventory         []string
	InventoryCap      int
	InventoryUpgrades int
	Money             int
	XP                int
	XPMax             int
	ManaMax           int
	Mana              int
	Skills            []string
	StoryProgress     int
}

func initCharacter(name string, class string) Character {
	return Character{
		Name:              name,
		Class:             class,
		Level:             1,
		BaseHpMax:         100,
		HpMax:             100,
		Hp:                100,
		AttackBase:        5,
		Inventory:         []string{},
		InventoryCap:      10,
		InventoryUpgrades: 0,
		Money:             100,
		XP:                0,
		XPMax:             20,
		ManaMax:           10,
		Mana:              10,
		Skills:            []string{},
		StoryProgress:     1,
	}
}
