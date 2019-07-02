package charts

import (
	"github.com/teawel/app/dbs"
	"github.com/teawel/app/utils"
	"log"
)

type LineChart struct {
	AxisChart `yaml:",inline" json:",inline"`

	Series []*LineSeries `yaml:"series" json:"series"`
	Labels []string      `yaml:"labels" json:"labels"`
}

func NewLineChart() *LineChart {
	return &LineChart{
		Series: []*LineSeries{},
		Labels: []string{},
	}
}

func (this *LineChart) AddSeries(series ...*LineSeries) {
	this.Series = append(this.Series, series...)
}

func (this *LineChart) Type() string {
	return TypeLine
}

func (this *LineChart) Fetch(db *dbs.DB) error {
	for lineIndex, line := range this.Series {
		values, err := db.Query([]string{utils.TimeFormat("Ymd")}, line.Query)
		if err != nil {
			log.Println("[error]" + err.Error())
			continue
		}
		if values == nil {
			continue
		}
		valuesSlice, ok := values.([]map[string]interface{})
		if !ok {
			continue
		}
		for _, v := range valuesSlice {
			if lineIndex == 0 {
				this.Labels = append(this.Labels, v["label"].(string))
			}
			line.AddValue(v["value"], v["label"].(string))
		}
	}
	return nil
}
