package app

import (
	"github.com/teawel/app/charts"
	"github.com/teawel/app/dbs"
	"github.com/teawel/app/utils"
)

const (
	CanvasQuarter = 0.25 // 1/4
	CanvasHalf    = 0.5  // 1/2
	CanvasFull    = 1    // 1/1
)

// chart canvas
type ChartCanvas struct {
	Id          string                `yaml:"id" json:"id"`
	Name        string                `yaml:"name" json:"name"`
	Type        string                `yaml:"type" json:"type"`
	OptionsJSON string                `yaml:"options" json:"options"`
	Chart       charts.ChartInterface `yaml:"chart" json:"chart"`

	WidthPercent  float32 `yaml:"widthPercent" json:"widthPercent"`
	HeightPercent float32 `yaml:"heightPercent" json:"heightPercent"`

	LeftMenu  *charts.Menu `yaml:"leftMenu" json:"leftMenu"`
	RightMenu *charts.Menu `yaml:"rightMenu" json:"rightMenu"`
}

func NewChartCanvas(id string, name string, widthPercent float32, heightPercent float32) *ChartCanvas {
	return &ChartCanvas{
		Id:            id,
		Name:          name,
		WidthPercent:  widthPercent,
		HeightPercent: heightPercent,
	}
}

func (this *ChartCanvas) SetChart(chart charts.ChartInterface) {
	this.Type = chart.Type()
	this.OptionsJSON = string(utils.JSONEncode(chart))
}

func (this *ChartCanvas) SetParam(param *charts.Param) {
	if param == nil {
		return
	}
	if param.MenuParam != nil && param.MenuParam.HasItems() {
		if this.LeftMenu != nil {
			this.LeftMenu.SelectItem(param.MenuParam.ItemIds ...)
		}
		if this.RightMenu != nil {
			this.RightMenu.SelectItem(param.MenuParam.ItemIds ...)
		}
	}
	if this.Chart != nil {
		this.Chart.SetParam(param)
	}
}

func (this *ChartCanvas) Fetch(db *dbs.DB) error {
	chart, err := charts.Decode(this.Type, this.OptionsJSON)
	if err != nil {
		return err
	}
	err = chart.Fetch(db)
	if err != nil {
		return err
	}

	this.Chart = chart
	return nil
}
