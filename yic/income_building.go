package yic

import "math"

type incomeBuildingEffect struct {
	level int
}

func (b *incomeBuildingEffect) IncomePerSecond() float64 {
	return 15.0 * math.Pow10(b.level)
}

func (b *incomeBuildingEffect) Upgradable() bool {
	return b.level < 3
}

func (b *incomeBuildingEffect) Upgrade() {
	if b.level < 3 {
		b.level += 1
	}
}

func (b *incomeBuildingEffect) Cost() float64 {
	return 5 * math.Pow10(b.level+2)
}

func (b *incomeBuildingEffect) Level() int {
	return b.level
}
