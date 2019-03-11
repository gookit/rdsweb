package api

import "github.com/gookit/rux"

const RdsName = `/{name:[\w-]+}`

// RedisAPI controller
type RedisAPI struct {
	BaseAPI
}

// AddRoutes for the controller
func (a *RedisAPI) AddRoutes(g *rux.Router) {
	g.GET(RdsName+"/open", a.Open)
	g.GET(RdsName+"/info", a.Info)
	g.GET(RdsName+"/stats", a.Stats)
	g.GET(RdsName+"/search", a.SearchKey)
	g.GET(RdsName+"/list-db", a.ListDB)
	g.GET(RdsName+"/list-keys", a.ListKeys)
}

// Open
func (a *RedisAPI) Open(c *rux.Context) {

}

// Info
func (a *RedisAPI) Info(c *rux.Context) {

}

// Stats
func (a *RedisAPI) Stats(c *rux.Context) {

}

// ListDB
func (a *RedisAPI) ListDB(c *rux.Context) {

}

// ListKeys
func (a *RedisAPI) ListKeys(c *rux.Context) {

}

// SearchKey
func (a *RedisAPI) SearchKey(c *rux.Context) {

}
