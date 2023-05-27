package Extensions

// NeighboursCount Расчёт количества соседей
func NeighboursCount(area [][]bool, row, column int) byte {
	var neighbours byte

	for idRow := -1; idRow <= 1; idRow++ {
		for idColumn := -1; idColumn <= 1; idColumn++ {

			idNeighRow, idNeighCol := row+idRow, column+idColumn

			if idRow == 0 && idColumn == 0 {
				continue
			}
			if idNeighRow < 0 || idNeighCol < 0 {
				continue
			}
			if idNeighRow >= len(area) || idNeighCol >= len(area[0]) {
				continue
			}

			if area[idNeighRow][idNeighCol] {
				neighbours++
			}
		}
	}

	return neighbours
}
