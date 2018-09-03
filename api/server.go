package api

import (
	"github.com/gookit/redis-viewer/app"
	"github.com/gookit/respond"
	"github.com/gookit/sux"
)

type ServerAPI struct {
}

// AddRoutes for the controller
func (a *ServerAPI) AddRoutes(g *sux.Router) {
	g.GET("", a.Index)
	g.GET("/{name}", a.Get)
}

// Servers get redis server names from config
func (a *ServerAPI) Index(c *sux.Context) {
	ss, _ := app.Cfg.Strings("servers", ",")

	respond.JSON(c.Resp, 200, ss)
}

func (a *ServerAPI) Get(c *sux.Context) {

}

func ListDB(c *sux.Context) {

}
