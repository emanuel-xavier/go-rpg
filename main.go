package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	PlayerImage *ebiten.Image
	x, y        float64
	Speed       float64
}

func (p *Player) MoveUp() { p.y -= p.Speed }

func (p *Player) MoveDown() { p.y += p.Speed }

func (p *Player) MoveRight() { p.x += p.Speed }

func (p *Player) MoveLeft() { p.x -= p.Speed }

type Game struct {
	Player Player
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.Player.MoveUp()
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Player.MoveDown()
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.Player.MoveLeft()
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.Player.MoveRight()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.Player.x, g.Player.y)

	// draw the player
	screen.DrawImage(
		g.Player.PlayerImage.SubImage(
			image.Rect(0, 0, 16, 16),
		).(*ebiten.Image),
		&opts,
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeithg int) {
	return 320, 240
	// return ebiten.WindowSize()
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	playerImg, _, err := ebitenutil.NewImageFromFile("assets/images/Noble/SpriteSheet.png")
	if err != nil {
		log.Fatal(err)
	}

	game := &Game{
		Player: Player{
			x:           0,
			y:           0,
			Speed:       2,
			PlayerImage: playerImg,
		},
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
