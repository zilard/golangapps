package main

import "fmt"

type Position struct {
	x float64
	y float64
}

func (p *Position) Move(x, y float64) {
	p.x += x
	p.y += y
}

func (p *Position) Teleport(x, y float64) {
	p.x = x
	p.y = y
}

type Player struct {
	*Position
}

func NewPlayer() *Player {
	return &Player{
		Position: &Position{},
	}
}

type Enemy struct {
	posX float64
	posY float64
}

func (e *Enemy) Move(x, y float64) {
	e.posX += x
	e.posY += y
}

func main() {
	player := NewPlayer()

	fmt.Println(player.Position)
}
