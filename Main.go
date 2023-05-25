package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const (
	windowWidth  = 640
	windowHeight = 480
)

type Game struct {
	pixels []byte
	world  *World
}

// Update Обновление данных на каждой итерации
func (g *Game) Update() error {
	g.world.Update()
	return nil
}

// Draw Отрисовка данных на каждой итерации
func (g *Game) Draw(screen *ebiten.Image) {

	// Создается двумерный массив с пикселями
	if g.pixels == nil {
		g.pixels = make([]byte, windowWidth*windowHeight*4)
	}

	// Заполняем массив пикселей данными об ареи игры
	g.world.Draw(g.pixels)

	// Считывается полученный массив пикселей и выводится на экран
	screen.WritePixels(g.pixels)
}

// Layout инициализирует размеры поля
func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return windowWidth, windowHeight
}

// Инициализация на старте
func init() {

}

func main() {

	game := &Game{
		world: NewWorld(windowWidth, windowHeight),
	}

	ebiten.SetWindowSize(windowWidth*2, windowHeight*2)
	ebiten.SetWindowTitle("Game of life")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
