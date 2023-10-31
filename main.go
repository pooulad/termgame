package main

import (
	"bytes"
	"fmt"
)

const (
	player  = 69
	wall    = 1
	nothing = 0
)

type level struct {
	width, height int
	data          [][]byte
}

func newLevel(width, height int) *level {
	data := make([][]byte, 18)
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			data[h] = make([]byte, width)
		}
	}
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			if w == width-1 {
				data[h][w] = wall
			}
			if w == 0 {
				data[h][w] = wall
			}
			if h == 0 {
				data[h][w] = wall
			}
			if h == height-1 {
				data[h][w] = wall
			}
		}
	}
	return &level{
		data: data,
	}
}
type game struct {
	isRunning bool
}

func newGame(width, height int) *game {
	return &game{}
}

func (g *game) start() {
	g.isRunning = true
	g.loop()
}
func (g *game) loop() {
	for g.isRunning {
		g.update()
		g.render()
	}
}
func (g *game) update() {}
func (g *game) render() {}

func main() {
	width := 80
	height := 18
	level := make([][]byte, 18)

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			level[h] = make([]byte, width)
		}
	}
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			if w == width-1 {
				level[h][w] = wall
			}
			if w == 0 {
				level[h][w] = wall
			}
			if h == 0 {
				level[h][w] = wall
			}
			if h == height-1 {
				level[h][w] = wall
			}
		}
	}
	buf := new(bytes.Buffer)
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			if level[h][w] == nothing {
				buf.WriteString(" ")
			}
			if level[h][w] == wall {
				buf.WriteString("H")
			}
		}
		buf.WriteString("\n")
	}
	fmt.Println(buf.String())
}
