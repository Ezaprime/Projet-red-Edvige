package piscine

type Monster struct {
	Name   string
	HpMax  int
	Hp     int
	Attack int
}

func Fight(c *Character, m *Monster) {
	m.Hp -= c.AttackBase
	if m.Hp < 0 {
		m.Hp = 0
	}
}
