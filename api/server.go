package api

import (
	"github.com/gookit/ini"
	"github.com/gookit/rdsweb/app"
	"github.com/gookit/rux"
)

// BaseAPI controller
type BaseAPI struct {
}

// ServerAPI controller
type ServerAPI struct {
	BaseAPI
}

// AddRoutes for the controller
func (a *ServerAPI) AddRoutes(g *rux.Router) {
	g.GET("", a.Index)
	g.GET("/names", a.Names)
	g.GET(RdsName, a.Get)
	g.POST("", a.Create)
	g.PUT(RdsName, a.Update)
	g.DELETE(RdsName, a.Delete)
}

// Servers get redis server list
func (a *ServerAPI) Index(c *rux.Context) {
	var servers []ini.Section
	for _, name := range app.Names {
		conf := ini.StringMap(name)
		if len(conf) > 0 {
			conf["name"] = name
			servers = append(servers, conf)
		}
	}

	c.JSONBytes(200, app.JSON(servers))
}

// Names get redis server names from config
func (a *ServerAPI) Names(c *rux.Context) {
	ss := ini.Strings("servers", ",")

	c.JSONBytes(200, app.JSON(ss))
}

// Get a redis server config by name
func (a *ServerAPI) Get(c *rux.Context) {
	name := c.Param("name")
	conf := ini.StringMap(name)

	if len(conf) > 0 {
		conf["name"] = name
		c.JSONBytes(200, app.JSON(conf))
	} else {
		c.JSONBytes(404, app.ErrJSON(404, "not found"))
	}
}

// Create new redis server config
func (a *ServerAPI) Create(c *rux.Context) {

}

// Update a redis server config
func (a *ServerAPI) Update(c *rux.Context) {

}

// Delete a redis server config
func (a *ServerAPI) Delete(c *rux.Context) {
	name, ok := c.QueryParam("name")
	if !ok || name == "" {
		bs := app.ErrJSON(2, "invalid request params")
		c.JSONBytes(200, bs)
		return
	}

	ini.Default().DelSection(name)
	c.JSONBytes(200, app.JSON(nil))
}
