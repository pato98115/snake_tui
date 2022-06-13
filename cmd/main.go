package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	game "github.com/pato98115/snake_tui/pkg/game"
)

// "██"

const (
	backgroundStr string = "██"
	food1Str      string = "██"
	food2Str      string = "██"
	food3Str      string = "██"
	snakeBodyStr  string = "██"
	snakeHeadStr  string = "██"

	backgroundColor string = "#9A9A9A" // grey
	food1Color      string = "#F9B4FD" // ligth pink
	food2Color      string = "#FB80C5" // pink
	food3Color      string = "#B51E73" // strong pink
	snakeBodyColor  string = "#2BEC54" // green
	snakeHeadColor  string = "#29E050" // other green
)

var baseStyle = lipgloss.NewStyle().
	Bold(true).UnsetAlign().UnsetMargins().UnsetPadding()

func strCellType(cell game.CellType) (string, error) {
	var strCell, strColor string
	switch cell {
	case game.BackgroundCell:
		strCell = backgroundStr
		strColor = backgroundColor
	case game.FoodCell1:
		strCell = food1Str
		strColor = food1Color
	case game.FoodCell2:
		strCell = food2Str
		strColor = food2Color
	case game.FoodCell3:
		strCell = food3Str
		strColor = food3Color
	case game.SnakeBodyCell:
		strCell = snakeBodyStr
		strColor = snakeBodyColor
	case game.SnakeHeadCell:
		strCell = snakeHeadStr
		strColor = snakeHeadColor
	default:
		return "", fmt.Errorf("Error: invalid cell type %v", cell)
	}
	strCell = baseStyle.Foreground(lipgloss.Color(strColor)).Render(strCell)

	return strCell, nil
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

type tickMsg struct{}

type model struct {
	game     *game.Game
	stepTime time.Duration
}

func tick(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}

func (m model) Init() tea.Cmd {
	return tick(m.stepTime)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Make sure these keys always quit
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			return m, tea.Quit
		}
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "w": // up
			m.game.ChangeDir(game.Up)
		case "d": // right
			m.game.ChangeDir(game.Right)
		case "s": // down
			m.game.ChangeDir(game.Down)
		case "a": // left
			m.game.ChangeDir(game.Left)
		}
		return m, nil
	case tickMsg:
		running := m.game.Step()
		if !running {
			return m, tea.Quit
		}

		return m, tick(m.stepTime)
	}

	return m, nil
}

func (m model) View() string {
	var strRepr string
	var err error
	var typeRepr []game.CellType

	typeRepr = m.game.Represent()

	strRepr, err = strReprBoard(
		m.game.Board.Cols,
		m.game.Board.Rows,
		typeRepr,
	)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	return strRepr
}

func main() {
	initialModel := model{
		game:     game.New(),
		stepTime: time.Second / 2,
	}

	p := tea.NewProgram(initialModel)
	if err := p.Start(); err != nil {
		fmt.Println("could not start program:", err)
	}
}
