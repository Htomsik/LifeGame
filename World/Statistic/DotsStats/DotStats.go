package DotsStats

import "LifeGame/World/Statistic/Extensions"

type DotStats struct {
	Count  int
	Coords [][2]int
}

// Init	Инициализация статистики умирающих точке
func (dots *DotStats) Init() {
	dots.Coords = make([][2]int, 0)
}

func (dots *DotStats) Next(area [][]bool) {
	dots.Init()
	dots.detectDots(area)
	dots.Count = len(dots.Coords) * 2
}

// detectDots обнаружение точек
func (dots *DotStats) detectDots(area [][]bool) {

	for idRow, column := range area {
		for idColumn, field := range column {

			neighbours := Extensions.NeighboursCount(area, idRow, idColumn)

			if neighbours > 3 || neighbours < 2 && field {
				dots.AddDot(idRow, idColumn)
			}
		}
	}
}

// AddDot Добавление координат точки
func (dots *DotStats) AddDot(row, column int) {
	dotCoords := [2]int{row, column}

	dots.Coords = append(dots.Coords, dotCoords)
}
