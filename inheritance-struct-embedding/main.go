package main

type Player struct {
	posX float64
	posY float64
}

func (p *Player) Move(x, y float64) {
	p.posX += x
	p.posY += y
}

func (p *Player) Teleport(x, y float64) {
	p.posX = x
	p.posY = y
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

}
