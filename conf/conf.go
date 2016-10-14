package conf

import (
	"github.com/go-ini/ini"
	"log"
)

var Cfg *ini.File

func init() {
	NewConf()
}
func NewConf() *ini.File {
	var err error
	Cfg, err = ini.Load("./ini/admin.ini")
	if err != nil {
		log.Println("[Hm-Debug]: admin.ini load failed!")
	}
	return Cfg
}
