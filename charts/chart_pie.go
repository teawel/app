package charts

import (
	"github.com/teawel/app/dbs"
)

type PieChart struct {
	BasicChart

	Series []*PieSeries `yaml:"series" json:"series"`
}

func NewPieChart() *PieChart {
	return &PieChart{
		Series: []*PieSeries{},
	}
}

func (this *PieChart) Type() string {
	return TypePie
}

func (this *PieChart) AddSeries(series ...*PieSeries) {
	this.Series = append(this.Series, series...)
}

func (this *PieChart) Fetch(db *dbs.DB) error {
	for _, s := range this.Series {
		err := s.Fetch(db)
		if err != nil {
			return err
		}
	}
	return nil
}
