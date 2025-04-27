package app

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/shalldie/leek/internal/store"
)

type TableModel struct {
	table   table.Model
	getRows func() []table.Row
}

func (g TableModel) Init() tea.Cmd {

	return tea.Batch(nil)
}

func (t TableModel) Update(msg tea.Msg) (TableModel, tea.Cmd) {
	var cmd tea.Cmd
	switch msg.(type) {
	case store.CMD_UPDATE:
		// 更新内容、高度
		rows := t.getRows()
		// rows[0][1] = lipgloss.NewStyle().Foreground(COLOR_PINK).Render(rows[0][1])
		t.table.SetRows(rows)
		t.table.SetHeight(len(rows) + 2)
		return t, nil
	}

	return t, cmd
}

func (t TableModel) View() string {
	return baseStyle.Render(t.table.View())
}

func NewTable(columns []table.Column, getRows func() []table.Row) TableModel {
	t := table.New(
		table.WithColumns(columns),
		// table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(1),
		// table.WithHeight(len(rows)+1),
		// table.style
		// table.WithStyleFunc(func(row, col int, value string) lipgloss.Style {
		// 	if row == 5 {
		// 		if col == 1 {
		// 			return s.Cell.Copy().Background(lipgloss.Color("#006341"))
		// 		} else if col == 2 {
		// 			return s.Cell.Copy().Background(lipgloss.Color("#FFFFFF"))
		// 		} else if col == 3 {
		// 			return s.Cell.Copy().Background(lipgloss.Color("#C8102E"))
		// 		} else {
		// 			return s.Cell
		// 		}
		// 	}

		// 	return s.Cell
		// }),
	)

	s := table.DefaultStyles()

	// s.Cell = s.Cell.Bold(true).Foreground(COLOR_PINK)

	s.Header = s.Header.
		BorderStyle(lipgloss.ASCIIBorder()).
		Foreground(COLOR_BLUE).
		// BorderForeground(COLOR_BLUE).
		BorderForeground(baseStyle.GetBorderTopForeground()).
		// BorderForeground(lipgloss.Color("201")).
		BorderBottom(true).
		Bold(true)
	s.Selected = s.Selected.
		// Foreground(lipgloss.Color("229")).
		// Background(lipgloss.Color("57")).
		UnsetForeground().
		Bold(false)
	t.SetStyles(s)

	return TableModel{
		table:   t,
		getRows: getRows,
	}
}
