package main

import (
	"LifeGame/World"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"math/rand"
	"time"
)

const (
	windowWidth  = 640
	windowHeight = 480
)

type Game struct {
	pixels []byte
	world  *World.World
}

// Update Обновление данных на каждой итерации
func (game *Game) Update() error {
	game.world.Next()
	return nil
}

// Draw Отрисовка данных на каждой итерации
func (game *Game) Draw(screen *ebiten.Image) {

	// Создается двумерный массив с пикселями
	if game.pixels == nil {
		game.pixels = make([]byte, windowWidth*windowHeight*4)
	}

	// Заполняем массив пикселей данными об ареи игры
	game.world.Draw(game.pixels)

	// Считывается полученный массив пикселей и выводится на экран
	screen.WritePixels(game.pixels)

	// Вывод тпс
	msg := fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS())
	ebitenutil.DebugPrint(screen, msg)
}

// Layout инициализирует размеры поля
func (game *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return windowWidth, windowHeight
}

// Инициализация на старте
func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	game := &Game{
		world: World.NewWorld(windowWidth, windowHeight, windowWidth*windowHeight/2),
	}

	ebiten.SetWindowSize(windowWidth*2, windowHeight*2)
	ebiten.SetWindowTitle("Game of life")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
