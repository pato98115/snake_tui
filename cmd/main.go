package main

import (
	"fmt"

	game "github.com/pato98115/snake_tui/pkg/game"
)

const (
	backgroundStr string = "[B]"
	food1Str      string = "[1]"
	food2Str      string = "[2]"
	food3Str      string = "[3]"
	snakeBodyStr  string = "[S]"
	snakeHeadStr  string = "[H]"
)

func strCellType(cell game.CellType) (string, error) {
	switch cell {
	case game.BackgroundCell:
		return backgroundStr, nil
	case game.FoodCell1:
		return food1Str, nil
	case game.FoodCell2:
		return food2Str, nil
	case game.FoodCell3:
		return food3Str, nil
	case game.SnakeBodyCell:
		return snakeBodyStr, nil
	case game.SnakeHeadCell:
		return snakeHeadStr, nil
	default:
		return "", fmt.Errorf("Error: invalid cell type %v", cell)
	}
}

func strReprBoard(cols, rows uint, cells []game.CellType) (string, error) {
	repr := ""
	for i := uint(0); i < rows; i++ {
		for j := uint(0); j < cols; j++ {
			reprCell, err := strCellType(cells[(i*cols)+j])

			if err != nil {
				return "", err
			}
			repr += reprCell
		}
		repr += "\n"
	}

	return repr, nil
}

func main() {
	var strRepr string
	var err error
	var typeRepr []game.CellType

	g := game.New()

	typeRepr = g.Represent()

	strRepr, err = strReprBoard(g.Board.Cols, g.Board.Rows, typeRepr)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Printf(strRepr)

	fmt.Println("\nRunning %v\n", g.Step())

	typeRepr = g.Represent()

	strRepr, err = strReprBoard(g.Board.Cols, g.Board.Rows, typeRepr)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Printf(strRepr)

	g.ChangeDir(game.Down)

	fmt.Println("\nRunning %v\n", g.Step())
	fmt.Println("\nRunning %v\n", g.Step())
	fmt.Println("\nRunning %v\n", g.Step())
	fmt.Println("\nRunning %v\n", g.Step())
	fmt.Println("\nRunning %v\n", g.Step())

	typeRepr = g.Represent()

	strRepr, err = strReprBoard(g.Board.Cols, g.Board.Rows, typeRepr)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Printf(strRepr)

	fmt.Println("\nRunning %v\n", g.Step())

	typeRepr = g.Represent()

	strRepr, err = strReprBoard(g.Board.Cols, g.Board.Rows, typeRepr)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Printf(strRepr)

	g.ChangeDir(game.Right)

	fmt.Println("\nRunning %v, %v\n", g.Step())

	g.ChangeDir(game.Up)

	fmt.Println("\nRunning %v, %v\n", g.Step())

	g.ChangeDir(game.Left)

	fmt.Println("\nRunning %v, %v\n", g.Step())

	typeRepr = g.Represent()

	strRepr, err = strReprBoard(g.Board.Cols, g.Board.Rows, typeRepr)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Printf(strRepr)
}
