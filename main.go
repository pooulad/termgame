package main

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

const (
	PLAYER      = 69
	WALL        = 1
	NOTHING     = 0
	MAX_SAMPLES = 100
)

type input struct {
	pressedKey byte
}

func (i *input) update() {
	b := make([]byte, 0)
	os.Stdin.Read(b)
	i.pressedKey = b[0]
}

type position struct {
	x, y int
}

type player struct {
	pos     position
	level   *level
	input   *input
	reverce bool
}

func (p *player) update() {
	if p.reverce {
		p.pos.x -= 1
		if p.pos.x == 2 {
			p.pos.x += 1
			p.reverce = false
		}
		return
	}
	p.pos.x += 1
	if p.pos.x == p.level.width-2 {
		p.pos.x -= 1
		p.reverce = true
	}
}

type stats struct {
	start  time.Time
	frames int
	fps    float64
}

func newStats() *stats {
	return &stats{
		start: time.Now(),
		fps:   69,
	}
}

func (s *stats) update() {
	s.frames++
	if s.frames == MAX_SAMPLES {
		s.fps = float64(s.frames) / float64(time.Since(s.start).Seconds())
		s.frames = 0
		s.start = time.Now()
	}
}

type level struct {
	width, height int
	data          [][]int
}

func newLevel(width, height int) *level {
	data := make([][]int, 18)
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			data[h] = make([]int, width)
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

func (l *level) set(pos position, v int) {
	l.data[pos.y][pos.x] = v
}

type game struct {
	isRunning bool
	level     *level
	stats     *stats
	player    *player
	input     *input
	drawBuf   *bytes.Buffer
}

func newGame(width, height int) *game {
	lvl := newLevel(width, height)
	inp := &input{}
	return &game{
		level:   lvl,
		drawBuf: new(bytes.Buffer),
		stats:   newStats(),
		input:   &input{},
		player: &player{
			level: lvl,
			input: inp,
			pos: position{
				x: 2,
				y: 5,
			},
		},
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
		g.stats.update()
		time.Sleep(time.Millisecond * 16)
	}
}
func (g *game) update() {
	g.level.set(g.player.pos, NOTHING)
	g.player.update()
	g.level.set(g.player.pos, PLAYER)
}

// func (g *game) renderPlayer() {
// 	g.level.set(g.player.pos, 69)
// }

func (g *game) renderLevel() {
	for h := 0; h < g.level.height; h++ {
		for w := 0; w < g.level.width; w++ {
			if g.level.data[h][w] == NOTHING {
				g.drawBuf.WriteString(" ")
			}
			if g.level.data[h][w] == WALL {
				g.drawBuf.WriteString("□")
			}
			if g.level.data[h][w] == PLAYER {
				g.drawBuf.WriteString("♛")
			}
		}
		g.drawBuf.WriteString("\n")
	}
}

func (g *game) render() {
	g.drawBuf.Reset()
	fmt.Fprint(os.Stdout, "\033[2J\033[1;1H")
	g.renderLevel()
	// g.renderPlayer()
	g.renderStats()
	fmt.Fprint(os.Stdout, g.drawBuf.String())
}

func (g *game) renderStats() {
	g.drawBuf.WriteString("-- STATS\n")
	g.drawBuf.WriteString(fmt.Sprintf("FPS: %.2f", g.stats.fps))
}

func main() {
	width := 80
	height := 18

	g := newGame(width, height)
	g.start()
}
