package app

type OperationHandler func(options map[string]string) (result string, err error)

type Operation struct {
	Code        string `yaml:"code" json:"code"`
	Name        string `yaml:"name" json:"name"`
	Description string `yaml:"description" json:"description"`

	handler OperationHandler
}

func NewOperation() *Operation {
	return &Operation{}
}

func (this *Operation) OnRun(handler OperationHandler) {
	this.handler = handler
}

func (this *Operation) Handler() OperationHandler {
	return this.handler
}
