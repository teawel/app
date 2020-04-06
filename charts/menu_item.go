package charts

type MenuItem struct {
	Id       string `yaml:"id" json:"id"`
	Name     string `yaml:"name" json:"name"`
	IsActive bool   `yaml:"isActive" json:"isActive"`
}

func NewMenuItem(id, name string) *MenuItem {
	return &MenuItem{
		Id:   id,
		Name: name,
	}
}
