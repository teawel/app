package app

type CommandMsg struct {
	Id   string
	Code string
	Args map[string]string
}
