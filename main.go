package main

import (
	"go-rpg/entities"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Go rpg")
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

	tileMapJson, err := NewTileMapJSON("./assets/maps/spawn.json")
	if err != nil {
		log.Fatal(err)
	}

	tilemapImg, _, err := ebitenutil.NewImageFromFile("./assets/images/tilesets/TilesetFloor.png")
	if err != nil {
		log.Fatal(err)
	}

	game := &Game{
		TilemapJSON: tileMapJson,
		TilemapImg:  tilemapImg,
		Player: &entities.Player{
			Sprite: &entities.Sprite{
				X:   0,
				Y:   0,
				Img: playerImg,
			},
			Speed:  2,
			Health: 100,
		},
		Enemies: []*entities.Enemy{
			&entities.Enemy{
				Sprite: &entities.Sprite{
					X:   50,
					Y:   50,
					Img: pandaImg,
				},
				Speed: 1.5,
			},
			&entities.Enemy{
				FollowsPlayer: true,
				Sprite: &entities.Sprite{
					X:   100,
					Y:   50,
					Img: pandaImg,
				},
				Speed: 1.0,
			},
		},
		Potions: []*entities.Potion{
			&entities.Potion{
				Sprite: &entities.Sprite{
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
