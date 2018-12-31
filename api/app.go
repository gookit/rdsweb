package api

import (
	"github.com/gookit/redis-viewer/app"
	"github.com/gookit/respond"
	"github.com/gookit/sux"
)

func Home(c *rux.Context) {
	respond.JSON(c.Resp, 200, rux.M{"hello": "welcome"})
}

func Config(c *rux.Context) {
	respond.JSON(c.Resp, 200, app.Cfg.Data())
}
