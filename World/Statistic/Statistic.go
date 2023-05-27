package Statistic

import (
	"LifeGame/World/Statistic/DotsStats"
	"LifeGame/World/Statistic/SquareStats"
)

type Statistic struct {
	SquareStats.SquareStats
	DotsStats.DotStats
	Cycle int
}

// Next Обновление статистики
func (stats *Statistic) Next(area [][]bool) {
	stats.Cycle++

	stats.SquareStats.Next(area)
	stats.DotStats.Next(area)
}

// Init	Инициализация статистики
func (stats *Statistic) Init() {
	stats.SquareStats.Init()
}
