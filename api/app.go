package api

import (
	"github.com/gookit/redis-viewer/app"
	"github.com/gookit/respond"
	"github.com/gookit/sux"
)

func Home(c *sux.Context) {
	respond.JSON(c.Resp, 200, sux.M{"hello": "welcome"})
}

func Config(c *sux.Context) {
	respond.JSON(c.Resp, 200, app.Cfg.Data())
}
