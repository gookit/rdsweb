package app

import (
	"fmt"
	"github.com/gobuffalo/packr"
	"github.com/gookit/i18n"
	"github.com/gookit/ini"
	"github.com/gookit/respond"
	"github.com/json-iterator/go"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
	Cfg  *ini.Ini
	Res  packr.Box
)

// Boot app components
func Boot() {
	var err error

	// init config
	Cfg, err = ini.LoadFiles("conf/config.ini")
	if err != nil {
		panic(err)
	}

	// init resource pack
	Res = packr.NewBox("../resource")

	// init responder
	respond.Marshal = json.Marshal
	respond.MarshalIndent = json.MarshalIndent
	respond.Initialize(func(opts *respond.Options) {
		opts.TplLayout = "layout-2col.tpl"
		opts.TplViewsDir = "templates"
	})

	loadServerNames()

	loadLanguages()
}

func loadLanguages() {
	defI18n := i18n.Default()
	defI18n.DefaultLang = "en"
	defI18n.NewLang("en", "English")
	defI18n.NewLang("zh-CN", "简体中文")

	// load data
	for lang, _ := range defI18n.Languages() {
		str := Res.String(fmt.Sprintf("languages/%s.ini", lang))
		defI18n.Lang(lang).LoadStrings(str)
	}
}
