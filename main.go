package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct {
	Img  *ebiten.Image
	X, Y float64
}

type Player struct {
	*Sprite
	Speed float64
}

func (p *Player) MoveUp() { p.Y -= p.Speed }

func (p *Player) MoveDown() { p.Y += p.Speed }

func (p *Player) MoveRight() { p.X += p.Speed }

func (p *Player) MoveLeft() { p.X -= p.Speed }

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

type Potion struct {
	*Sprite
	AmtHeal float64
}

type Game struct {
	Player  Player
	Enemies []*Enemy
	Potions []*Potion
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

	for _, enemy := range g.Enemies {
		enemy.Move(g.Player.X, g.Player.Y)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.Player.X, g.Player.Y)

	// draw the player
	screen.DrawImage(
		g.Player.Img.SubImage(
			image.Rect(0, 0, 16, 16),
		).(*ebiten.Image),
		&opts,
	)

	// Draw others sprites
	opts.GeoM.Reset()
	for _, enemy := range g.Enemies {
		opts.GeoM.Translate(enemy.X, enemy.Y)
		screen.DrawImage(
			enemy.Img.SubImage(
				image.Rect(0, 0, 16, 16),
			).(*ebiten.Image),
			&opts,
		)

		opts.GeoM.Reset()
	}

	// Draw potions
	opts.GeoM.Reset()
	for _, potion := range g.Potions {
		opts.GeoM.Translate(potion.X, potion.Y)
		screen.DrawImage(
			potion.Img.SubImage(
				image.Rect(0, 0, 16, 16),
			).(*ebiten.Image),
			&opts,
		)

		opts.GeoM.Reset()
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeithg int) {
	return 320, 240
	// return ebiten.WindowSize()
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	playerImg, _, err := ebitenutil.NewImageFromFile("./assets/images/Noble/SpriteSheet.png")
	if err != nil {
		log.Fatal(err)
	}

	pandaImg, _, err := ebitenutil.NewImageFromFile("./assets/images/Panda/SpriteSheet.png")
	if err != nil {
		log.Fatal(err)
	}

	potionImg, _, err := ebitenutil.NewImageFromFile("./assets/images/Potion/LifePot.png")
	if err != nil {
		log.Fatal(err)
	}

	game := &Game{
		Player: Player{
			Sprite: &Sprite{
				X:   0,
				Y:   0,
				Img: playerImg,
			},
			Speed: 2,
		},
		Enemies: []*Enemy{
			&Enemy{
				Sprite: &Sprite{
					X:   50,
					Y:   50,
					Img: pandaImg,
				},
				Speed: 1.5,
			},
			&Enemy{
				FollowsPlayer: true,
				Sprite: &Sprite{
					X:   100,
					Y:   50,
					Img: pandaImg,
				},
				Speed: 1.0,
			},
		},
		Potions: []*Potion{
			&Potion{
				Sprite: &Sprite{
					X:   200,
					Y:   150,
					Img: potionImg,
				},
				AmtHeal: 1.0,
			},
		},
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
