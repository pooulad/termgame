package main

import "fmt"

const (
	player  = 69
	wall    = 1
	nothing = 0
)

type game struct {
}

func (g *game) update() {}
func (g *game) render() {}

func main() {
	width := 20
	height := 20
	level := make([][]byte, 20)

	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			level[h] = make([]byte, width)
		}
	}
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			level[h][w] = wall
		}
	}
	fmt.Println("test")
}
