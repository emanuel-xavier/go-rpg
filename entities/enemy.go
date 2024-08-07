package entities

type Enemy struct {
	*Sprite
	Speed         float64
	FollowsPlayer bool
}

func (e *Enemy) Move(playerX, playerY float64) {
	if !e.FollowsPlayer {
		return
	}
	if e.X < playerX {
		e.X += e.Speed
	} else if e.X > playerX {
		e.X -= e.Speed
	}
	if e.Y < playerY {
		e.Y += e.Speed
	} else if e.Y > playerY {
		e.Y -= e.Speed
	}

}
