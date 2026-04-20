package main

import (
	"image/color"
	"math/rand"

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
	ticker    int
}

func NewGame() *Game {
	return &Game{
		snake:     []Point{{X: 10, Y: 10}, {X: 9, Y: 10}, {X: 8, Y: 10}},
		direction: Point{X: 1, Y: 0},
		food:      Point{X: 20, Y: 20},
	}
}

func (g *Game) Update() error {
	g.ticker++
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyArrowUp):
		g.direction = Point{X: 0, Y: -1}
	case ebiten.IsKeyPressed(ebiten.KeyArrowDown):
		g.direction = Point{X: 0, Y: 1}
	case ebiten.IsKeyPressed(ebiten.KeyArrowLeft):
		g.direction = Point{X: -1, Y: 0}
	case ebiten.IsKeyPressed(ebiten.KeyArrowRight):
		g.direction = Point{X: 1, Y: 0}
	}
	// food: Point{X: rand.Intn(30), Y: rand.Intn(30)}
	if g.ticker%10 == 0 {
		head := g.snake[0]
		newHead := Point{
			X: head.X + g.direction.X,
			Y: head.Y + g.direction.Y,
		}

		if newHead == g.food {
			g.food = Point{X: rand.Intn(30), Y: rand.Intn(30)}
			g.snake = append([]Point{newHead}, g.snake...)
		} else {
			g.snake = append([]Point{newHead}, g.snake[:len(g.snake)-1]...)
		}
	}
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
	return 600, 600
}

func main() {
	ebiten.SetWindowSize(600, 600)
	ebiten.SetWindowTitle("Snake")
	ebiten.RunGame(NewGame())
}
