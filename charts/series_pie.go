package charts

import (
	"github.com/teawel/app/dbs"
	"github.com/teawel/app/utils"
)

type PieSeries struct {
	Series

	Radius string    `yaml:"radius" json:"radius"`
	Center *Position `yaml:"center" json:"center"`
}

func NewPieSeries() *PieSeries {
	return &PieSeries{
		Series: Series{
			Values: []*Value{},
		},
	}
}

func (this *PieSeries) Fetch(db *dbs.DB) error {
	if len(this.Query) == 0 {
		return nil
	}
	result, err := db.Query([]string{utils.TimeFormat("Ymd")}, this.Query)
	if err != nil {
		return err
	}
	valueSlice, ok := result.([]map[string]interface{})
	if !ok {
		return nil
	}
	if len(valueSlice) == 0 {
		return nil
	}
	lastValue := valueSlice[len(valueSlice)-1]
	value := lastValue["value"]
	label := lastValue["label"]
	subValues, ok := value.([]string)
	if !ok {
		return nil
	}
	subLabels, ok := label.([]string)
	if !ok {
		return nil
	}
	for index, v := range subValues {
		if index < len(subLabels) {
			this.AddValue(v, subLabels[index])
		} else {
			this.AddValue(v, "")
		}
	}
	return nil
}
