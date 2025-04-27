package app

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/shalldie/leek/internal/store"
)

var (
	app *tea.Program
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	// BorderForeground(lipgloss.Color("240")).
	BorderForeground(COLOR_BORDER)

type AppModel struct {
	gold TableModel
}

func (m AppModel) Init() tea.Cmd {

	return tea.Batch(nil)
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// case "esc":
		// 	if m.table.Focused() {
		// 		m.table.Blur()
		// 	} else {
		// 		m.table.Focus()
		// 	}
		case "q", "ctrl+c":
			return m, tea.Quit
		}

	case CMD_UPDATE:

		return m, nil
	}

	m.gold, cmd = m.gold.Update(msg)
	return m, cmd
}

func (m AppModel) View() string {
	// return baseStyle.Render(m.gold.View()) + "\n"
	return m.gold.View()
}

func Run() {
	// list := []*stock.Stock{
	// 	stock.NewAU9999(),
	// 	stock.NewAU0(),
	// 	stock.NewXAU(),
	// 	stock.NewGC(),
	// }

	// columns := []table.Column{
	// 	{Title: "名称", Width: 10},
	// 	{Title: "涨跌", Width: 8},
	// 	{Title: "涨幅", Width: 8},
	// }

	// rows := []table.Row{}

	// for _, item := range list {
	// 	// rate := lipgloss.NewStyle().Foreground(COLOR_PINK).Render(item.Rate)
	// 	rows = append(rows, []string{item.Name, item.Rise, item.Rate})
	// }

	// t := table.New(
	// 	table.WithColumns(columns),
	// 	table.WithRows(rows),
	// 	table.WithFocused(true),
	// 	table.WithHeight(len(rows)+1),
	// )

	// s := table.DefaultStyles()

	// s.Header = s.Header.
	// 	BorderStyle(lipgloss.ASCIIBorder()).
	// 	Foreground(COLOR_BLUE).
	// 	// BorderForeground(COLOR_BLUE).
	// 	BorderForeground(baseStyle.GetBorderTopForeground()).
	// 	// BorderForeground(lipgloss.Color("201")).
	// 	BorderBottom(true).
	// 	Bold(true)
	// s.Selected = s.Selected.
	// 	// Foreground(lipgloss.Color("229")).
	// 	// Background(lipgloss.Color("57")).
	// 	UnsetForeground().
	// 	Bold(false)
	// t.SetStyles(s)

	app = tea.NewProgram(AppModel{
		gold: NewTable([]table.Column{
			{Title: "名称", Width: 10},
			{Title: "涨跌", Width: 8},
			{Title: "涨幅", Width: 8},
		}, func() []table.Row {
			rows := []table.Row{}

			for _, item := range store.State.Golds {
				// rate := lipgloss.NewStyle().Foreground(COLOR_PINK).Render(item.Rate)
				rows = append(rows, []string{item.Name, item.Rise, item.Rate})
			}
			return rows
		}),
	})

	go func() {
		store.SendImpl = func(cmd any) {
			app.Send(cmd)
		}

		store.State.Reset()
		store.Send(store.CMD_UPDATE(""))
		go func() {
			time.Sleep(time.Second * 3)
			store.State.Update()
			store.Send(store.CMD_UPDATE(""))
		}()
	}()

	if _, err := app.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
