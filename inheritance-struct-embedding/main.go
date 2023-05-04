package main

import "fmt"

type SpecialPosition struct {
	Position
}

func (sp *SpecialPosition) MoveSpecial(x, y float64) {
	sp.x += x * x
	sp.y += y * y
}

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
	*SpecialPosition
}

func NewEnemy() *Enemy {
	return &Enemy{
		SpecialPosition: &SpecialPosition{},
	}
}

func main() {
	raidBoss := NewEnemy()
	raidBoss.Move(1.1, 10.4)
	fmt.Println("raidBoss:", raidBoss.Position)

	player := NewPlayer()
	player.Move(1.1, 10.4)
	fmt.Println(player.Position)
	player.Teleport(1000.4, 3000.4)
	fmt.Println(player.Position)
}
