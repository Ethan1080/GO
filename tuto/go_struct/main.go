package main

import "fmt"

type Player struct {
	X int
	Y int
}

func (p *Player) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

func main() {
	p := Player{0, 0}
	p.Move(2, 2)

	fmt.Println(p.X, p.Y)
}
