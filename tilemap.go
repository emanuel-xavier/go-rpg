package main

import (
	"encoding/json"
	"log"
	"os"
	"path"
)

type TilemapLayserJSON struct {
	Data   []int  `json:"data"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Name   string `json:"name"`
}

type TilemapJSON struct {
	Layers   []TilemapLayserJSON `json:"layers"`
	Tilesets []map[string]any    `json:"tilesets"`
}

func (t *TilemapJSON) GenTilesets() ([]Tileset, error) {
	tilesets := make([]Tileset, len(t.Tilesets))

	for i, tilesetData := range t.Tilesets {
		tilesetPath := path.Join("./assets/maps/", tilesetData["source"].(string))
		log.Println("tilesetPath", tilesetPath)
		tileset, err := NewTileset(tilesetPath, int(tilesetData["firstgid"].(float64)))
		if err != nil {
			log.Println("New tileset", err)
			return nil, err
		}

		tilesets[i] = tileset
	}

	return tilesets, nil
}

func NewTileMapJSON(filepath string) (*TilemapJSON, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Println("Reading file", err)
		return nil, err
	}

	var tilemapJSON TilemapJSON
	err = json.Unmarshal(content, &tilemapJSON)
	if err != nil {
		log.Println("Unmarshal timemapJSON", err)
		return nil, err
	}

	return &tilemapJSON, nil
}
