package charts

import "github.com/teawel/app/dbs"

type ChartInterface interface {
	Type() string
	Fetch(db *dbs.DB) error
}
