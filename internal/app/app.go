package app

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/shalldie/leek/internal/app/stock_list"
	"github.com/shalldie/leek/internal/store"
	"github.com/shalldie/leek/internal/utils"
)

var (
	app      *tea.Program
	interval *utils.IntervalTimer
)

type AppModel struct {
	stockList stock_list.StockListModel
	table     TableModel
}

func (m AppModel) Init() tea.Cmd {

	return tea.Batch(nil)
}

func (m AppModel) propagate(msg tea.Msg) (AppModel, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	m.table, cmd = m.table.Update(msg)
	cmds = append(cmds, cmd)

	m.stockList, cmd = m.stockList.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			interval.Stop()
			return m, tea.Quit
		}

		// case CMD_UPDATE:
		// 	return m, nil
	}

	return m.propagate(msg)
}

func (m AppModel) View() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		lipgloss.NewStyle().Padding(1, 2, 0, 1).Render(m.stockList.View()),
		m.table.View(),
	)
}

func Run() {

	app = tea.NewProgram(AppModel{
		table: NewTable([]string{"名称", "价格", "涨跌", "涨幅"}, func() [][]string {
			rows := [][]string{}

			for _, item := range store.State.Stocks() {
				rows = append(rows, []string{item.Name, item.Price, item.Rise, item.Rate})
			}
			return rows
		}),
		stockList: stock_list.NewStockListModel(),
	})

	go func() {
		store.SendImpl = func(cmd any) {
			app.Send(cmd)
		}

		store.State.Reset()
		store.Send(store.CMD_UPDATE(""))

		var handler = func() {
			store.State.Update()
			store.Send(store.CMD_UPDATE(""))
		}

		handler()
		interval = utils.NewIntervalTimer(time.Second, handler)
	}()

	if _, err := app.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
