package app

type Fetcher func(options map[string]string) (result map[string]string, err error)
