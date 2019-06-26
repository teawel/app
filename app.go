package app

type App struct {
	Id           string `yaml:"id" json:"id"` // distributed by wel.teaos.cn
	Name         string `yaml:"name" json:"name"`
	Description  string `yaml:"description" json:"description"`
	Developer    string `yaml:"developer" json:"developer"`
	Site         string `yaml:"site" json:"site"`
	Version      string `yaml:"version" json:"version"`
	DocumentSite string `yaml:"documentSite" json:"documentSite"`
	SourceSite   string `yaml:"sourceSite" json:"sourceSite"`
	DownloadSite string `yaml:"downloadSite" json:"downloadSite"`
}

func NewApp() *App {
	return &App{}
}
