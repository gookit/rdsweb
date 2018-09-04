package api

import "github.com/gookit/sux"

const RdsName = `/{name:[\w-]+}`

// RedisAPI controller
type RedisAPI struct {
	BaseAPI
}

// AddRoutes for the controller
func (a *RedisAPI) AddRoutes(g *sux.Router) {
	g.GET(RdsName+"/open", a.Open)
	g.GET(RdsName+"/info", a.Info)
	g.GET(RdsName+"/stats", a.Stats)
	g.GET(RdsName+"/search", a.SearchKey)
	g.GET(RdsName+"/list-db", a.ListDB)
	g.GET(RdsName+"/list-keys", a.ListKeys)
}

// Open
func (a *RedisAPI) Open(c *sux.Context) {

}

// Info
func (a *RedisAPI) Info(c *sux.Context) {

}

// Stats
func (a *RedisAPI) Stats(c *sux.Context) {

}

// ListDB
func (a *RedisAPI) ListDB(c *sux.Context) {

}

// ListKeys
func (a *RedisAPI) ListKeys(c *sux.Context) {

}

// SearchKey
func (a *RedisAPI) SearchKey(c *sux.Context) {

}
