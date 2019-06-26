package charts

import (
	"github.com/teawel/app/dbs"
	"github.com/teawel/app/utils"
	"log"
)

type LineChart struct {
	BasicChart

	Lines    []*Line     `yaml:"lines" json:"lines"`
	MaxValue interface{} `yaml:"maxValue" json:"maxValue"`
	Labels   []string    `yaml:"labels" json:"labels"`
}

func NewLineChart() *LineChart {
	return &LineChart{
		Lines: []*Line{},
	}
}

func (this *LineChart) AddLine(line ...*Line) {
	this.Lines = append(this.Lines, line...)
}

func (this *LineChart) Type() string {
	return TypeLine
}

func (this *LineChart) Fetch(db *dbs.DB) error {
	for lineIndex, line := range this.Lines {
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
