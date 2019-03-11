package main

import (
	"github.com/gobuffalo/packr"
	"github.com/gookit/rdsweb/api"
	"github.com/gookit/rdsweb/app"
	"github.com/gookit/respond"
	"github.com/gookit/rux"
)

func main() {
	app.Boot()

	r := rux.New()
	rux.Debug(true)

	addRoutes(r)
	r.Listen(":18080")
}

func addRoutes(r *rux.Router) {
	assets := packr.NewBox("./static")
	r.StaticFS("/static", assets)

	r.GET("/test", api.Home)
	r.GET("/conf", api.Config)

	r.Controller("/redis", &api.RedisAPI{})
	r.Controller("/servers", &api.ServerAPI{})

	r.GET("/", func(c *rux.Context) {
		respond.HTMLString(c.Resp, 200, app.Res.String("templates/index.html"), rux.M{
			"name": "inhere",
		})
	})
}
