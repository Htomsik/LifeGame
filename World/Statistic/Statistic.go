package Statistic

type Statistic struct {
	Cycle        int
	SquareCount  int
	SquareCoords [][2]int
}

// Next Обновление статистики
func (stats *Statistic) Next(area [][]bool) {
	stats.Cycle++

	stats.Init()
	stats.detectSquares(area)
}

// Init	Инициализация статистики
func (stats *Statistic) Init() {
	stats.SquareCoords = make([][2]int, 0)
}

// detectSquares обнаружение количества
// TODO Подумать над оптимизацией
func (stats *Statistic) detectSquares(area [][]bool) {

	for idRow, column := range area {
		for idColumn, _ := range column {

			if coords, isDetected := isSquare(area, idRow, idColumn); isDetected {
				for _, elem := range coords {
					stats.SquareCoords = append(stats.SquareCoords, elem)
				}
			}
		}
	}

	stats.SquareCount = len(stats.SquareCoords) / 4
}

// isSquare Проверяет входит ли поле в структуру квадрата
func isSquare(area [][]bool, row, column int) ([4][2]int, bool) {
	var squareSides [4][2]int
	var countSides int

	for idRow := 0; idRow >= -1; idRow-- {
		for idColumn := 0; idColumn >= -1; idColumn-- {

			checkedRow := row + idRow
			checkedCol := column + idColumn

			if checkedRow < 0 || checkedCol < 0 {
				continue
			}

			if area[checkedRow][checkedCol] {

				squareSides[countSides][0] = checkedRow
				squareSides[countSides][1] = checkedCol

				countSides++
			} else {
				countSides = 0
				break
			}
		}
	}

	return squareSides, countSides == 4
}
