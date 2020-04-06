package charts

import "github.com/teawel/app/lists"

type Menu struct {
	Items []*MenuItem `yaml:"items" json:"items"`
}

func NewMenu() *Menu {
	return &Menu{
		Items: []*MenuItem{},
	}
}

func (this *Menu) AddItem(item *MenuItem) {
	this.Items = append(this.Items, item)
}

func (this *Menu) SelectItem(itemId ...string) {
	if len(itemId) == 0 {
		return
	}
	for _, item := range this.Items {
		if lists.ContainsString(itemId, item.Id) {
			item.IsActive = true
		} else {
			item.IsActive = false
		}
	}
}
