package app

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/shalldie/leek/internal/store"
	"github.com/shalldie/leek/internal/utils"
)

var (
	app *tea.Program
)

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
	return m.gold.View()
}

func Run() {

	app = tea.NewProgram(AppModel{
		gold: NewTable([]string{"名称", "价格", "涨跌", "涨幅"}, func() [][]string {
			rows := [][]string{}

			for _, item := range store.State.Golds {
				rows = append(rows, []string{item.Name, item.Price, item.Rise, item.Rate})
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

		var handler = func() {
			store.State.Update()
			store.Send(store.CMD_UPDATE(""))
		}

		handler()
		utils.NewIntervalTimer(time.Second, handler)
	}()

	if _, err := app.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
