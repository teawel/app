package charts

import (
	"github.com/teawel/app/dbs"
	"github.com/teawel/app/types"
	"github.com/teawel/app/utils"
)

type ValueChart struct {
	BasicChart `yaml:",inline" json:",inline"`

	Value      *Value `yaml:"value" json:"value"`
	ValueColor string `yaml:"valueColor" json:"valueColor"` // value color
	LabelColor string `yaml:"labelColor" json:"labelColor"` // label color
	Query      string `yaml:"query" json:"query"`
}

func NewValueChart() *ValueChart {
	return &ValueChart{}
}

func (this *ValueChart) Type() string {
	return TypeValue
}

func (this *ValueChart) Fetch(db *dbs.DB) error {
	if len(this.Query) > 0 {
		result, err := db.Query([]string{utils.TimeFormat("Ymd")}, this.Query)
		if err != nil {
			return err
		}
		if result == nil {
			return nil
		}
		valuesSlice, ok := result.([]map[string]interface{})
		if !ok {
			return nil
		}
		value := valuesSlice[len(valuesSlice)-1]
		this.Value = &Value{
			Value: value["value"],
			Label: types.String(value["label"]),
		}
	}
	return nil
}
