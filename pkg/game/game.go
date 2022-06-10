package game

import (
	"fmt"

	deque "github.com/pato98115/snake_tui/pkg/deque"
)

// Snake direction in board
type direction byte

// Cell for composing snake's body
type cell struct {
	x uint
	y uint
}

// Snake
type snake struct {
	food uint
	body *deque.Deque
}

// Snake's food
type food struct {
	cell     *cell
	quantity uint
}

// Matrix like board
type board struct {
	rows uint
	cols uint
}

// Game
type Game struct {
	snake     snake
	snake_dir direction
	food      []food
	board     board
	points    uint
}

const (
	defaultRows uint = 40
	defaultCols uint = 40

	snakeStartSize uint = 3

	dirUp direction = iota
	dirDown
	dirLeft
	dirRight
)

// Compare two cells
func (c *cell) equals(c2 *cell) bool {
	return (c.x == c2.x && c.y == c2.y)
}

func (s *snake) move(c *cell) {
	s.body.PushFront(c)
	// has food inside to grow?
	if s.food > 0 {
		s.food -= 1
	} else {
		s.body.PopBack()
	}
}

func (s *snake) getHead() *cell {
	v := s.body.Front().Value
	return v.(*cell)
}

func (s *snake) eat(f food) error {
	if !s.getHead().equals(f.cell) {
		return fmt.Errorf("ERROR: Snake can't reach that")
	}
	s.food += f.quantity
	return nil
}

func (g *Game) changeDir(d direction) {
	// only change direction when it's a valid one
	// depending on current direction
	switch d {
	case dirUp, dirDown:
		switch g.snake_dir {
		case dirLeft, dirRight:
			g.snake_dir = d
		}
	case dirLeft, dirRight:
		switch g.snake_dir {
		case dirUp, dirDown:
			g.snake_dir = d
		}
	}
}
