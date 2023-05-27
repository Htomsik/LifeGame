package World

import (
	"LifeGame/World/Statistic"
	"LifeGame/World/Statistic/Extensions"
	"math/rand"
)

// RGBA цвета
const (
	White = 255
)

var (
	Blue = [4]byte{0, 200, 255, 1}
	Red  = [4]byte{255, 0, 0, 1}
)

// World Игровое поле
type World struct {
	Statistic.Statistic

	area         [][]bool
	width        int
	height       int
	TickDuration int
}

// NewWorld Создание нового мира
func NewWorld(width, height, initialFields, tickDuration int) *World {

	world := &World{
		area:         makeArea(width, height),
		width:        width,
		height:       height,
		TickDuration: tickDuration,
	}

	if initialFields > 0 {
		world.init(initialFields)
	}

	world.Statistic.Init()

	return world
}

// makeArea Создание нового пустого игрового мира
func makeArea(width, height int) [][]bool {
	area := make([][]bool, width)

	for id := range area {
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

			neighbours := Extensions.NeighboursCount(world.area, idRow, idColumn)

			if neighbours == 3 || neighbours == 2 && elem {
				newArea[idRow][idColumn] = true
			}
		}
	}

	world.area = newArea

	world.Statistic.Next(world.area)
}

// Draw Транслейт игрового поля в массив пикселей
func (world *World) Draw(pixels []byte) {

	for idRow := range world.area {
		for idColumn := range world.area[idRow] {

			isColored := world.area[idRow][idColumn]

			if !isColored {
				continue
			}

			setPixel(pixels, idColumn*world.width+idRow, White)
		}
	}

	for _, coords := range world.SquareStats.Coords {
		setPixelAll(pixels, coords[1]*world.width+coords[0], Blue)
	}

	for _, coords := range world.DotStats.Coords {
		setPixelAll(pixels, coords[1]*world.width+coords[0], Red)
	}
}

// setPixel Установка цвета пикселя
func setPixel(pixels []byte, id int, color byte) {
	for i := 0; i <= 3; i++ {
		pixels[4*id+i] = color
	}
}

// setPixel Установка цвета пикселя
func setPixelAll(pixels []byte, id int, colors [4]byte) {
	for i := 0; i <= 3; i++ {
		pixels[4*id+i] = colors[i]
	}
}
