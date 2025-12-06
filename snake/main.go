package main //salut salut

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	grid   [32][32]int
	snakeX int
	snakeY int
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyZ) {

	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS) {

	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyQ) {

	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD) {

	}

	if ebiten.IsKeyPressed(ebiten.KeyR) {
		g.reset()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	rows := len(g.grid)
	cols := len(g.grid[0])

	// Calculer la taille de chaque carré pour que ça remplisse la fenêtre
	cellWidth := float64(1920) / float64(cols)
	cellHeight := float64(1080) / float64(rows)

	for x := 0; x < 32; x++ {
		for y := 0; y < 32; y++ {
			if g.grid[x][y] == 1 {
				rect := ebiten.NewImage(int(cellWidth), int(cellHeight))
				rect.Fill(color.RGBA{255, 0, 0, 255}) // rouge
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x)*cellWidth, float64(y)*cellHeight)
				screen.DrawImage(rect, op)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1920, 1080
}

func (g *Game) reset() {
	g.grid = [32][32]int{}
	g.grid[1][1] = 1
}

func main() {
	fmt.Print("start")

	game := &Game{}
	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowTitle("snake")

	game.reset()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
