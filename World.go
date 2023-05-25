package main

const (
	White = 255
	Black = 0
)

// World Игровое поле
type World struct {
	area          []bool
	width, height int
}

// NewWorld Создание нового мира
func NewWorld(width, height int) *World {

	world := &World{
		area:   make([]bool, width*height),
		width:  width,
		height: height,
	}

	return world
}

// Update Обновление игрового поля
func (world *World) Update() {

}

// Draw Транслейт игрового поля в массив пикселей
func (world *World) Draw(pixels []byte) {
	for id, elem := range world.area {
		if elem {
			setPixel(pixels, id, White)
		} else {
			setPixel(pixels, id, Black)
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
