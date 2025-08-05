package stock_list

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/shalldie/leek/internal/store"
)

type StockListModel struct {
	list list.Model
}

func (s StockListModel) Init() tea.Cmd {
	return nil
}

func (m StockListModel) propagate(msg tea.Msg) (StockListModel, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	lastIndex := m.list.Index()
	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)
	curIndex := m.list.Index()
	if lastIndex != curIndex {
		go func() {
			store.State.MarketIndex = m.list.Index()
			store.Send(store.CMD_UPDATE(""))
		}()
	}

	m.refreshList()

	return m, tea.Batch(cmds...)
}

func (s StockListModel) Update(msg tea.Msg) (StockListModel, tea.Cmd) {

	return s.propagate(msg)
}

func (s StockListModel) View() string {

	return s.list.View()
}

func (s *StockListModel) refreshList() {
	items := []list.Item{}

	for _, market := range store.State.Markets {
		items = append(items, ListItem{Name: market.Name})
	}

	s.list.SetItems(items)
	s.list.SetWidth(1500)
	s.list.SetHeight(len(items) + 5)
}

func NewStockListModel() StockListModel {
	// listDelegate，list item 选中色
	listDelegate := list.NewDefaultDelegate()
	listDelegate.ShowDescription = false
	listDelegate.SetSpacing(0)
	listDelegate.Styles.SelectedTitle = listDelegate.Styles.SelectedTitle.
		Foreground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#00acf8"}).
		BorderStyle(lipgloss.ThickBorder()).
		BorderLeftForeground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#00acf8"}).
		Bold(true)

		// listDelegate.Styles.SelectedDesc = listDelegate.Styles.SelectedDesc.
		// 	Foreground(listDelegate.Styles.NormalDesc.GetForeground()).
		// 	BorderStyle(lipgloss.ThickBorder()).
		// 	BorderLeftForeground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#00acf8"})

	// listDelegate.Styles.SelectedDesc = listDelegate.Styles.
	// 	SelectedDesc.
	// 	BorderStyle(lipgloss.HiddenBorder())

	// list
	list := list.New([]list.Item{}, listDelegate, 0, 0)

	// list.Title = "Leek"
	list.SetShowTitle(false)
	list.DisableQuitKeybindings()
	list.SetShowHelp(false)
	list.SetShowPagination(false)
	list.SetShowStatusBar(false)
	list.SetShowFilter(false)
	list.SetFilteringEnabled(false)

	// list.KeyMap = newListKeyMap()
	// list.AdditionalFullHelpKeys = additionalKeyMap
	// list.FilterInput.Prompt = i18n.Get(i18nTpl, "filelist_filter")

	m := StockListModel{
		list: list,
	}
	// m.list.StartSpinner()
	m.refreshList()
	m.list.ToggleSpinner()
	return m
}
