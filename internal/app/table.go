package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/shalldie/leek/internal/store"
)

type TableModel struct {
	table *table.Table
	// table   table.Model
	getRows func() [][]string
}

func (g TableModel) Init() tea.Cmd {
	return tea.Batch(nil)
}

func (t TableModel) Update(msg tea.Msg) (TableModel, tea.Cmd) {
	var cmd tea.Cmd
	switch msg.(type) {
	case store.CMD_UPDATE:
		// 更新内容
		t.table.ClearRows()
		t.table.Rows(t.getRows()...)
		return t, nil
	}

	return t, cmd
}

func (t TableModel) View() string {
	return t.table.String()
	// return baseStyle.Render(t.table.View())
}

func NewTable(headers []string, getRows func() [][]string) TableModel {

	headerStyle := lipgloss.NewStyle().Bold(false).Align(lipgloss.Center)
	cellStyle := lipgloss.NewStyle().Padding(0, 2)

	table := table.New().
		Border(lipgloss.RoundedBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(COLOR_PRIMARY)).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return headerStyle
			case col == 0: // name
				return cellStyle
			default: // 指标
				return cellStyle.Align(lipgloss.Center)
			}
		}).
		Headers(headers...)

	return TableModel{
		table:   table,
		getRows: getRows,
	}
}
