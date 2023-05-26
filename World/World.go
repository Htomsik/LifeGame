package World

import (
	"math/rand"
)

const (
	White = 255
	Black = 0
)

// World Игровое поле
type World struct {
	area          [][]bool
	width, height int
}

// NewWorld Создание нового мира
func NewWorld(width, height, initialFields int) *World {

	world := &World{
		area:   makeArea(width, height),
		width:  width,
		height: height,
	}

	if initialFields > 0 {
		world.init(initialFields)
	}

	return world
}

// makeArea Создание нового пустого игрового мира
func makeArea(width, height int) [][]bool {
	area := make([][]bool, width)

	for id, _ := range area {
		area[id] = make([]bool, height)
	}

	return area
}

// init Инициализация рандомных полей
func (world *World) init(countPixels int) {

	for i := 0; i < countPixels; i++ {

		w := rand.Intn(world.width)
		h := rand.Intn(world.height)

		world.area[w][h] = true
	}
}

// Next Обновление игрового поля
func (world *World) Next() {

	newArea := makeArea(world.width, world.height)

	for idRow, row := range world.area {
		for idColumn, elem := range row {

			neighbours := world.neighboursCount(idRow, idColumn)

			switch {
			case neighbours > 3 || neighbours < 2:
				newArea[idRow][idColumn] = false

			case neighbours == 3 || neighbours == 2 && elem:
				newArea[idRow][idColumn] = true
			}
		}
	}

	world.area = newArea
}

// Draw Транслейт игрового поля в массив пикселей
func (world *World) Draw(pixels []byte) {

	for idRow, _ := range world.area {
		for idColumn, _ := range world.area[idRow] {

			if world.area[idRow][idColumn] {
				setPixel(pixels, idColumn*world.width+idRow, White)
			} else {
				setPixel(pixels, idColumn*world.width+idRow, Black)
			}
		}
	}

}

// setPixel Установка цвета пикселя
func setPixel(pixels []byte, id int, color byte) {
	pixels[4*id] = color
	pixels[4*id+1] = color
	pixels[4*id+2] = color
	pixels[4*id+3] = color
}

// neighboursCount Расчёт количества соседей
func (world *World) neighboursCount(row, column int) int {
	var neighbours int

	for idRow := -1; idRow <= 1; idRow++ {
		for idColumn := -1; idColumn <= 1; idColumn++ {

			idNeightRow, idNeightCol := row+idRow, column+idColumn

			if idNeightRow >= world.width || idNeightRow < 0 {
				continue
			}
			if idNeightCol >= world.height || idNeightCol < 0 {
				continue
			}
			if idRow == 0 && idColumn == 0 {
				continue
			}

			if world.area[idNeightRow][idNeightCol] {
				neighbours++
			}
		}
	}

	return neighbours
}