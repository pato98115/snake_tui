package game

import (
	"math/rand"

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
	defaultRows uint = 10
	defaultCols uint = 10

	snakeStartSize uint = 3

	Up Direction = iota
	Right
	Down
	Left

	BackgroundCell CellType = iota
	FoodCell1
	FoodCell2
	FoodCell3
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

func randFood(coln, rown, quantityn uint) *food {
	return &food{
		cell: &cell{
			x: uint(rand.Intn(int(coln))),
			y: uint(rand.Intn(int(rown))),
		},
		quantity: uint(rand.Intn(int(quantityn-1))) + 1,
	}
}

func foodType(f *food) CellType {
	switch f.quantity {
	case 1:
		return FoodCell1
	case 2:
		return FoodCell2
	default:
		return FoodCell3
	}
}

func New() *Game {
	snake := snake{
		food: 0,
		body: deque.New(),
	}
	snake.body.PushFront(&cell{0, 0})
	snake.body.PushFront(&cell{0, 1})
	snake.body.PushFront(&cell{0, 2})

	board := &Board{
		Cols:  defaultCols,
		Rows:  defaultRows,
		cells: make([]CellType, defaultCols*defaultRows),
	}

	food := deque.New()
	food.PushFront(randFood(board.Cols, board.Rows, 4))

	return &Game{
		snake:    snake,
		snakeDir: Right,
		food:     food,
		running:  true,
		Board:    *board,
		Points:   0,
	}
}

func (g *Game) ChangeDir(d Direction) {
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

func (g *Game) validMove(next *cell) bool {
	// check collisions with borders.
	// using uint 0 - 1 = max uint for easier check.
	if next.x > g.Board.Cols || next.y > g.Board.Rows {
		return false
	}
	// check collisions with snake body.
	for e := g.snake.body.Front().Next(); e != nil; e = e.Next() {
		c := e.Value.(*cell)
		if next.equals(c) {
			return false
		}
	}
	return true
}

func (g *Game) Step() bool {
	if !g.running {
		return false
	}

	// gen next cell for snake
	nextCell := cell{}
	snakeHead := g.snake.getHead()
	switch g.snakeDir {
	case Up:
		nextCell.x = snakeHead.x
		nextCell.y = snakeHead.y - 1
	case Left:
		nextCell.x = snakeHead.x - 1
		nextCell.y = snakeHead.y
	case Down:
		nextCell.x = snakeHead.x
		nextCell.y = snakeHead.y + 1
	case Right:
		nextCell.x = snakeHead.x + 1
		nextCell.y = snakeHead.y
	}

	if !g.validMove(&nextCell) {
		g.running = false
		return g.running
	}

	for e := g.food.Front(); e != nil; e = e.Next() {
		f := e.Value.(*food)
		if nextCell.equals(f.cell) {
			g.snake.eat(f)
			g.Points += f.quantity
			g.food.Remove(e)
			// Spawn new food
			g.food.PushFront(randFood(g.Board.Cols, g.Board.Rows, 4))

			break
		}
	}

	g.snake.move(&nextCell)

	return true
}

func (g *Game) Represent() []CellType {
	rows := g.Board.Rows
	cols := g.Board.Cols

	// set background
	for i := uint(0); i < rows; i++ {
		for j := uint(0); j < cols; j++ {
			g.Board.cells[(i*cols)+j] = BackgroundCell
		}
	}
	// set food
	for e := g.food.Front(); e != nil; e = e.Next() {
		f := e.Value.(*food)
		g.Board.cells[(f.cell.y*cols)+f.cell.x] = foodType(f)
	}
	// set snake head
	h := g.snake.getHead()
	g.Board.cells[(h.y*cols)+h.x] = SnakeHeadCell

	for e := g.snake.body.Front().Next(); e != nil; e = e.Next() {
		c := e.Value.(*cell)
		g.Board.cells[(c.y*cols)+c.x] = SnakeBodyCell
	}

	return g.Board.cells
}
