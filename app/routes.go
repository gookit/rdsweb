package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr"
	"net/http"
)

func AddRoutes(r *gin.Engine)  {
	tpls := packr.NewBox("./templates")
	assets := packr.NewBox("./static")

	r.StaticFS("/static", assets)
	r.StaticFS("/templates", tpls)

	// r.StaticFS("/more_static", http.Dir("my_file_system"))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/templates/index.html", gin.H{
			"Foo": "World",
		})
	})
}
