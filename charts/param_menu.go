package charts

import "github.com/teawel/app/lists"

type MenuParam struct {
	ItemIds []string `yaml:"itemIds" json:"itemIds"`
}

func (this *MenuParam) HasItems() bool {
	return len(this.ItemIds) > 0
}

func (this *MenuParam) IsActive(itemId string) bool {
	return lists.ContainsString(this.ItemIds, itemId)
}
