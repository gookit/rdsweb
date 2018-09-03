package app

import (
	"github.com/gookit/ini"
	"github.com/gookit/respond"
	"github.com/json-iterator/go"
)

var Cfg *ini.Ini

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

// Boot app components
func Boot() {
	var err error

	// init config
	Cfg, err = ini.LoadFiles("conf/config.ini")
	if err != nil {
		panic(err)
	}

	// init responder
	respond.Marshal = json.Marshal
	respond.MarshalIndent = json.MarshalIndent
	respond.Initialize(func(opts *respond.Options) {
		opts.TplLayout = "layout-2col.tpl"
		opts.TplViewsDir = "templates"
	})

	loadServerNames()
}

