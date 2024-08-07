package main

import (
	"go-rpg/entities"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Player      *entities.Player
	Enemies     []*entities.Enemy
	Potions     []*entities.Potion
	TilemapJSON *TilemapJSON
	TilemapImg  *ebiten.Image
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

	for _, layer := range g.TilemapJSON.Layers {
		for index, id := range layer.Data {
			x := index % layer.Width
			y := index / layer.Width

			x *= 16
			y *= 16

			srcX := (id - 1) % 22
			srcY := (id - 1) / 22

			opts.GeoM.Translate(float64(x), float64(y))

			screen.DrawImage(
				g.TilemapImg.SubImage(image.Rect(srcX, srcY, srcX+16, srcY+16)).(*ebiten.Image),
				&opts,
			)

			opts.GeoM.Reset()
		}
	}

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

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
	// return ebiten.WindowSize()
}
