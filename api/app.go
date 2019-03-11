package api

import (
	"github.com/gookit/ini"
	"github.com/gookit/respond"
	"github.com/gookit/rux"
)

func Home(c *rux.Context) {
	err := respond.JSON(c.Resp, 200, rux.M{"hello": "welcome"})
	c.AddError(err)
}

func Config(c *rux.Context) {
	respond.JSON(c.Resp, 200, ini.Data())
}
