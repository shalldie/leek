package stock_list

type ListItem struct {
	Name string
}

func (item ListItem) Title() string {
	return item.Name
}

func (item ListItem) Description() string {
	// return zone.Mark(item.ID+"des", withEllipsis(item.Content))
	return ""
}

func (item ListItem) FilterValue() string { return item.Name }
