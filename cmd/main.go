package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	game "github.com/pato98115/snake_tui/pkg/game"
)

// "██"

const (
	backgroundStr string = "  "
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
	Bold(true).
	UnsetAlign().
	UnsetMargins().
	UnsetPadding()

var boderStyle = lipgloss.NewStyle().
	Border(lipgloss.DoubleBorder(), true)

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

type keyMap struct {
	Up    key.Binding
	Right key.Binding
	Down  key.Binding
	Left  key.Binding
	Help  key.Binding
	Quit  key.Binding
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the keyMap interface.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// keyMap interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Left, k.Right}, // first column
		{k.Help, k.Quit},                // second column
	}
}

type model struct {
	game     *game.Game
	stepTime time.Duration
	keys     *keyMap
	help     help.Model
}

var defaultKeyMap = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "w"),
		key.WithHelp("↑/w", "move up "),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "s"),
		key.WithHelp("↓/s", "move down "),
	),
	Left: key.NewBinding(
		key.WithKeys("left", "a"),
		key.WithHelp("←/a", "move left "),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "d"),
		key.WithHelp("→/d", "move right "),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help "),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit "),
	),
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
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// If we set a width on the help menu it can it can gracefully truncate
		// its view as needed.
		m.help.Width = msg.Width

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Up):
			m.game.ChangeDir(game.Up)
		case key.Matches(msg, m.keys.Down):
			m.game.ChangeDir(game.Down)
		case key.Matches(msg, m.keys.Left):
			m.game.ChangeDir(game.Left)
		case key.Matches(msg, m.keys.Right):
			m.game.ChangeDir(game.Right)
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}

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

	helpView := m.help.View(m.keys)

	pointsRepr := fmt.Sprintf("Points: %d", m.game.Points)

	return boderStyle.Render(strRepr+pointsRepr) + "\n" + helpView
}

func main() {
	initialModel := model{
		game:     game.New(),
		stepTime: time.Second / 3,
		keys:     &defaultKeyMap,
	}

	p := tea.NewProgram(initialModel)
	if err := p.Start(); err != nil {
		fmt.Println("could not start program:", err)
	}
}
