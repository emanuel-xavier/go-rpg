package main

import (
	"go-rpg/entities"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SCREEN_WIDTH  = 320.0
	SCREEN_HEIGHT = 240.0
)

type Game struct {
	Player      *entities.Player
	Enemies     []*entities.Enemy
	Potions     []*entities.Potion
	TilemapJSON *TilemapJSON
	TileSets    []Tileset
	TilemapImg  *ebiten.Image
	Cam         *Camera
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

	g.Cam.FollowTarget(g.Player.X, g.Player.Y, SCREEN_WIDTH, SCREEN_HEIGHT)

	g.Cam.Constrain(
		float64(g.TilemapJSON.Layers[0].Width*16),
		float64(g.TilemapJSON.Layers[0].Height*16),
		SCREEN_WIDTH,
		SCREEN_HEIGHT,
	)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

	opts := ebiten.DrawImageOptions{}

	for layserIndex, layer := range g.TilemapJSON.Layers {
		for index, id := range layer.Data {
			if id == 0 {
				continue
			}

			x := (index % layer.Width) * 16
			y := (index / layer.Width) * 16

			img := g.TileSets[layserIndex].Img(id)

			opts.GeoM.Translate(float64(x), float64(y))

			opts.GeoM.Translate(0.0, -float64(img.Bounds().Dy()+16))

			opts.GeoM.Translate(g.Cam.X, g.Cam.Y)

			screen.DrawImage(img, &opts)

			opts.GeoM.Reset()
		}
	}

	opts.GeoM.Translate(g.Player.X, g.Player.Y)
	opts.GeoM.Translate(g.Cam.X+8, g.Cam.Y+8)

	// draw the player
	screen.DrawImage(
		g.Player.Img.SubImage(
			image.Rect(0, 0, 16, 16),
		).(*ebiten.Image),
		&opts,
	)

	// Draw enemies
	opts.GeoM.Reset()
	for _, enemy := range g.Enemies {
		enemyX := enemy.X
		enemyY := enemy.Y

		if enemyX+16 >= -g.Cam.X && enemyX <= -g.Cam.X+SCREEN_WIDTH && enemyY+16 >= -g.Cam.Y && enemyY <= -g.Cam.Y+SCREEN_HEIGHT {
			opts.GeoM.Translate(enemyX, enemyY)
			opts.GeoM.Translate(g.Cam.X, g.Cam.Y)

			screen.DrawImage(
				enemy.Img.SubImage(
					image.Rect(0, 0, 16, 16),
				).(*ebiten.Image),
				&opts,
			)

			opts.GeoM.Reset()
		}
	}

	// Draw potions
	opts.GeoM.Reset()
	for _, potion := range g.Potions {
		potionX := potion.X
		potionY := potion.Y

		if potionX+16 >= -g.Cam.X && potionX <= -g.Cam.X+SCREEN_WIDTH && potionY+16 >= -g.Cam.Y && potionY <= -g.Cam.Y+SCREEN_HEIGHT {
			opts.GeoM.Translate(potionX, potionY)
			opts.GeoM.Translate(g.Cam.X, g.Cam.Y)

			screen.DrawImage(
				potion.Img.SubImage(
					image.Rect(0, 0, 16, 16),
				).(*ebiten.Image),
				&opts,
			)

			opts.GeoM.Reset()
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
	// return ebiten.WindowSize()
}
