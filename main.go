package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Point struct {
	X, Y int
}

type Game struct {
	snake     []Point
	direction Point
	food      Point
}

func NewGame() *Game {
	return &Game{
		snake:     []Point{{X: 10, Y: 10}},
		direction: Point{X: 1, Y: 0},
		food:      Point{X: 20, Y: 20},
	}
}

func (g *Game) Update() error {
	head := g.snake[0]
	newHead := Point{
		X: head.X + g.direction.X,
		Y: head.Y + g.direction.Y,
	}
	g.snake = append(g.snake, newHead)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	for _, p := range g.snake {
		ebitenutil.DrawRect(screen, float64(p.X*20), float64(p.Y*20), 18, 18, color.RGBA{0, 255, 0, 255})
	}

	ebitenutil.DrawRect(screen, float64(g.food.X*20), float64(g.food.Y*20), 18, 18, color.RGBA{255, 0, 0, 255})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 620, 240
}

func main() {
	ebiten.RunGame(NewGame())
}
