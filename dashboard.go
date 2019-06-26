package app

type Dashboard struct {
	Id      string         `yaml:"id" json:"id"`
	Version string         `yaml:"version" json:"version"`
	Name    string         `yaml:"name" json:"name"`
	Charts  []*ChartCanvas `yaml:"charts" json:"charts"`
}

func NewDashboard(id, version, name string) *Dashboard {
	return &Dashboard{
		Id:      id,
		Version: version,
		Name:    name,
	}
}

func NewDefaultDashboard(version string) *Dashboard {
	return NewDashboard("default", version, "Default")
}

func (this *Dashboard) AddChart(chart *ChartCanvas) {
	if chart == nil {
		return
	}
	this.Charts = append(this.Charts, chart)
}
