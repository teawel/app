package app

import (
	"github.com/teawel/app/utils"
	"log"
	"os"
	"path/filepath"
)

var Root = ""

func init() {
	if utils.IsTesting() {
		pwd, err := os.Getwd()
		if err != nil {
			log.Println("[error]" + err.Error())
		} else {
			Root = pwd
		}
	} else {
		exe, err := os.Executable()
		if err != nil {
			log.Println("[error]" + err.Error())
		} else {
			Root = filepath.Dir(exe)
		}
	}
}
