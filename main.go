package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/redis-viewer/app"
)

func main() {
	// r := gin.New()
	r := gin.Default()

	// t, err := loadTemplate()
	// if err != nil {
	// 	panic(err)
	// }
	// r.SetHTMLTemplate(t)

	app.AddRoutes(r)

	r.Run(":18080")
}
