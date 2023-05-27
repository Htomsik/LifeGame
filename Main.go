package main

import (
	"LifeGame/World"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

const (
	windowWidth  = 640
	windowHeight = 480
	tickDuration = 1
)

var (
	tickCount int
)

type Game struct {
	pixels []byte
	world  *World.World
}

// Update Обновление данных на каждой итерации
func (game *Game) Update() error {
	tickCount++

	if tickCount == tickDuration {
		game.world.Next()
		tickCount = 0
	}
	return nil
}

// Draw Отрисовка данных на каждой итерации
func (game *Game) Draw(screen *ebiten.Image) {

	// Создается двумерный массив с пикселями
	game.pixels = make([]byte, windowWidth*windowHeight*4)

	// Заполняем массив пикселей данными об ареи игры
	game.world.Draw(game.pixels)

	// Считывается полученный массив пикселей и выводится на экран
	screen.WritePixels(game.pixels)

	// Вывод статистики
	msg := fmt.Sprintf("TPS: %v\nCycle: %v\nSqare count: %v", ebiten.ActualTPS(), game.world.Cycle, game.world.SquareCount)

	ebitenutil.DebugPrint(screen, msg)
}

// Layout инициализирует размеры поля
func (game *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return windowWidth, windowHeight
}

// Инициализация на старте
func init() {
	//rand.Seed(time.Now().UnixNano())
}

func main() {

	game := &Game{
		world: World.NewWorld(windowWidth, windowHeight, windowWidth*windowHeight/2, tickDuration),
	}

	ebiten.SetWindowSize(windowWidth*2, windowHeight*2)
	ebiten.SetWindowTitle("Game of life")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
