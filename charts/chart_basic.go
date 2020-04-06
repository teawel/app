package charts

import "github.com/teawel/app/dbs"

type BasicChart struct {
	param *Param
}

func (this *BasicChart) SetParam(param *Param) {
	this.param = param
}

func (this *BasicChart) Param() *Param {
	return this.param
}

func (this *BasicChart) Fetch(db *dbs.DB) error {
	return nil
}
