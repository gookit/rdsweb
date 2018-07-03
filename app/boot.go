package app

import (
	"gopkg.in/ini.v1"
)

var Cfg *ini.File

func init() {
	Cfg,_ = ini.LooseLoad("conf/config.ini")
}

func Redis()  {

}
