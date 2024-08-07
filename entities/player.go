package entities

type Player struct {
	*Sprite
	Health uint
	Speed  float64
}

func (p *Player) MoveUp() { p.Y -= p.Speed }

func (p *Player) MoveDown() { p.Y += p.Speed }

func (p *Player) MoveRight() { p.X += p.Speed }

func (p *Player) MoveLeft() { p.X -= p.Speed }
