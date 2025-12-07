package main //salut salut

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type point struct {
	x int
	y int
}

type player struct {
	headx int
	heady int
	dir   int
	body  point
}

type Game struct {
	grid  [32][32]int
	snake player
}

func (g *Game) Update() error {
	g.handleKey()
	g.moveSnakeHead()

	if g.snake.headx < 0 || g.snake.headx > 31 || g.snake.heady > 32 || g.snake.heady < 0 {
		fmt.Println("Perdu!")
		//perdu g.reset()
		for {
			//
		}
	}

	g.UpdateGrid()

	//time.Sleep(1 * time.Second)

	return nil
}

func (g *Game) UpdateGrid() {
	g.grid = [32][32]int{}
	g.grid[g.snake.headx][g.snake.heady] = 1
}

func (g *Game) moveSnakeHead() {
	switch g.snake.dir {
	case 1:
		g.snake.heady += 1
	case 2:
		g.snake.heady += -1
	case 3:
		g.snake.headx += -1
	case 4:
		g.snake.headx += 1
	}
}

func (g *Game) handleKey() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyZ) {
		g.snake.dir = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		g.snake.dir = 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyQ) {
		g.snake.dir = 3
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		g.snake.dir = 4
	}

	if ebiten.IsKeyPressed(ebiten.KeyR) {
		g.reset()
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	rows := len(g.grid)
	cols := len(g.grid[0])

	cellWidth := float64(920) / float64(cols)
	cellHeight := float64(920) / float64(rows)

	for x := 0; x < 32; x++ {
		for y := 0; y < 32; y++ {
			if g.grid[x][y] == 1 {
				rect := ebiten.NewImage(int(cellWidth), int(cellHeight))
				rect.Fill(color.RGBA{255, 0, 0, 255})
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x)*cellWidth, float64(y)*cellHeight)
				screen.DrawImage(rect, op)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 920, 920
}

func (g *Game) reset() {
	g.grid = [32][32]int{}
	g.snake.headx = 3
	g.snake.headx = 1
}

func main() {
	fmt.Println("start")

	game := &Game{}
	ebiten.SetWindowSize(920, 920)
	ebiten.SetWindowTitle("snake")

	game.reset()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
