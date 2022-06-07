package game

import "github.com/gammazero/deque"

const (
	defaultRows = 40
	defaultCols = 40
)

type cell struct {
	x int
	y int
}

type snake struct {
	food int
	body *deque.Deque[cell]
}

func (s *snake) move(c cell) {
	s.body.PushBack(c)
	// has food inside to grow?
	if s.food > 0 {
		s.food -= 1
	} else {
		s.body.PopFront()
	}
}

type food struct {
	cell
	quantity int
}

type board struct {
	rows int
	cols int
}

type game struct {
	snake  snake
	food   []food
	board  board
	points int
}

func newGame() *game {
	return &game{
		snake: snake{
			food: 0,
			body: deque.New[cell](),
		},
		food: []food{},
		board: board{
			rows: defaultRows,
			cols: defaultCols,
		},
		points: 0,
	}
}
