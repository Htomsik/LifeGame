package SquareStats

type SquareStats struct {
	Count  int
	Coords [][2]int
}

// Init	Инициализация статистики квадратов
func (squares *SquareStats) Init() {
	squares.Coords = make([][2]int, 0)
}

// Next Обновление статистики
func (squares *SquareStats) Next(area [][]bool) {
	squares.Init()
	squares.detectSquares(area)
	squares.Count = len(squares.Coords) / 4
}

// detectSquares обнаружение квадратов
// TODO Подумать над оптимизацией
func (squares *SquareStats) detectSquares(area [][]bool) {

	for idRow, column := range area {
		for idColumn, _ := range column {

			if coords, isDetected := isSquare(area, idRow, idColumn); isDetected {
				for _, elem := range coords {
					squares.Coords = append(squares.Coords, elem)
				}
			}
		}
	}
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
