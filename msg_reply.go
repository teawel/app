package app

type ReplyMsg struct {
	Id     string
	Code   string
	Result interface{}
	Error  string
}
