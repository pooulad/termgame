package main

import (
	"bytes"
	"fmt"
	"os"
)

const (
	PLAYER      = 69
	WALL        = 1
	NOTHING     = 0
	MAX_SAMPLES = 100
)

type stats struct {
	start  time.Time
	frames int
	fps    float64
}

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
				data[h][w] = WALL
			}
			if w == 0 {
				data[h][w] = WALL
			}
			if h == 0 {
				data[h][w] = WALL
			}
			if h == height-1 {
				data[h][w] = WALL
			}
		}
	}
	return &level{
		data:   data,
		width:  width,
		height: height,
	}
}
type game struct {
	isRunning bool
	level     *level
	drawBuf   *bytes.Buffer
}

func newGame(width, height int) *game {
	lvl := newLevel(width, height)
	return &game{
		level:   lvl,
		drawBuf: new(bytes.Buffer),
	}
}

func (g *game) start() {
	g.isRunning = true
	g.loop()
}

func (g *game) loop() {
	for g.isRunning {
		g.update()
		g.render()
		time.Sleep(time.Millisecond * 13)
	}
}
func (g *game) update() {}

func (g *game) renderLevel() {
	for h := 0; h < g.level.height; h++ {
		for w := 0; w < g.level.width; w++ {
			if g.level.data[h][w] == NOTHING {
				g.drawBuf.WriteString(" ")
			}
			if g.level.data[h][w] == WALL {
				g.drawBuf.WriteString("□")
			}
		}
		g.drawBuf.WriteString("\n")
	}
}

func (g *game) render() {
	g.drawBuf.Reset()
	fmt.Fprint(os.Stdout, "\033[2J\033[1;1H")
	g.renderLevel()
	g.renderStats()
	fmt.Fprint(os.Stdout, g.drawBuf.String())
}

func (g *game) renderStats() {
	g.drawBuf.WriteString("-- STATS\n")
	g.drawBuf.WriteString(fmt.Sprintf("FPS: %.2f", 3.3))
}

func main() {
	width := 80
	height := 18

	g := newGame(width, height)
	g.start()
}
