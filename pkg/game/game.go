package game

import (
	deque "github.com/pato98115/snake_tui/pkg/deque"
)

// Snake direction in board
type Direction byte

// board cell type
type CellType byte

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
type Board struct {
	Rows  uint
	Cols  uint
	cells []CellType
}

// Game
type Game struct {
	snake    snake
	snakeDir Direction
	food     *deque.Deque
	Board    Board
	Points   uint
	running  bool
}

const (
	defaultRows uint = 40
	defaultCols uint = 40

	snakeStartSize uint = 3

	Up Direction = iota
	Right
	Down
	Left

	BackgroundCell CellType = iota
	FruitCell1
	FruitCell2
	FruitCell3
	SnakeBodyCell
	SnakeHeadCell
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

func (s *snake) eat(f *food) error {
	s.food += f.quantity
	return nil
}

func fruitType(f *food) CellType {
	switch f.quantity {
	case 1:
		return FruitCell1
	case 2:
		return FruitCell2
	default:
		return FruitCell3
	}
}

func (g *Game) changeDir(d Direction) {
	// only change direction when it's a valid one
	// depending on current direction
	switch d {
	case Up, Down:
		switch g.snakeDir {
		case Left, Right:
			g.snakeDir = d
		}
	case Left, Right:
		switch g.snakeDir {
		case Up, Down:
			g.snakeDir = d
		}
	}
}

func (g *Game) step() bool {
	// gen next cell for snake
	nextCell := cell{}
	snakeHead := g.snake.getHead()
	switch g.snakeDir {
	case Up:
		nextCell.x = snakeHead.x
		nextCell.y = snakeHead.y + 1
	case Right:
		nextCell.x = snakeHead.x + 1
		nextCell.y = snakeHead.y
	case Down:
		nextCell.x = snakeHead.x
		nextCell.y = snakeHead.y - 1
	case Left:
		nextCell.x = snakeHead.x - 1
		nextCell.y = snakeHead.y
	}

	// using uint 0 - 1 = max uint for easier check
	if nextCell.x > g.Board.Cols || nextCell.y > g.Board.Rows {
		g.running = false
		return g.running
	}

	for e := g.food.Front(); e != nil; e = e.Next() {
		f := e.Value.(*food)
		if nextCell.equals(f.cell) {
			g.snake.eat(f)
			g.Points += f.quantity
			g.food.Remove(e)
			break
		}
	}

	g.snake.move(&nextCell)

	return true
}

func (g *Game) buildBoard() []CellType {
	// set background
	for i := uint(0); i < g.Board.Rows; i++ {
		for j := uint(0); j < g.Board.Cols; j++ {
			g.Board.cells[i+j] = BackgroundCell
		}
	}
	// set fruits
	for e := g.food.Front(); e != nil; e = e.Next() {
		f := e.Value.(*food)
		g.Board.cells[f.cell.x+f.cell.y] = fruitType(f)
	}
	// set snake head
	h := g.snake.getHead()
	g.Board.cells[h.x+h.y] = SnakeHeadCell

	for e := g.snake.body.Front().Next(); e != nil; e = e.Next() {
		c := e.Value.(*cell)
		g.Board.cells[c.x+c.y] = SnakeBodyCell
	}

	return g.Board.cells
}
