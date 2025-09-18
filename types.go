package piscine

type Equipment struct {
	Head   string
	Torso  string
	Feet   string
	Weapon string
}

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
	Equip             Equipment
	Initiative        int
	XP                int
	XPMax             int
	ManaMax           int
	Mana              int
	Skills            []string
	StoryProgress     int
}

type Monster struct {
	Name       string
	HpMax      int
	Hp         int
	Attack     int
	Initiative int
}

type Chapter struct {
	Title        string
	Factory      func() Monster
	IsBoss       bool
	XpReward     int
	CreditReward int
}

var chapters = []Chapter{
	{"Rues de la Guill' : Police municipale", initMunicipale, false, 40, 60},
	{"Berges du Rhône : Police nationale", initNationale, false, 60, 80},
	{"Périph' : Gendarmerie", initGendarmerie, false, 80, 100},
	{"Tunnel : Motards", initMotards, false, 100, 120},
	{"Quartiers Nord : BAC", initBAC, false, 120, 150},
	{"Opération Finale : FBI", initBossFBI, true, 200, 250},
}
