package main

import (
	"github.com/gobuffalo/packr"
	"github.com/gookit/redis-viewer/api"
	"github.com/gookit/redis-viewer/app"
	"github.com/gookit/respond"
	"github.com/gookit/sux"
)

func main() {
	app.Boot()

	r := sux.New()
	sux.Debug(true)

	addRoutes(r)
	r.Listen(":18080")
}

func addRoutes(r *sux.Router) {
	assets := packr.NewBox("./static")
	r.StaticFS("/static", assets)

	tpls := packr.NewBox("./templates")

	r.GET("/test", api.Home)
	r.GET("/conf", api.Config)
	r.GET("/servers", api.Servers)

	r.Controller("/servers", &api.ServerAPI{})

	r.GET("/", func(c *sux.Context) {
		respond.HTMLString(c.Resp, 200, tpls.String("index.html"), sux.M{
			"name": "inhere",
		})
	})
}
