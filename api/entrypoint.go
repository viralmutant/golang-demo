package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var  (
	app *gin.Engine
)

func myRoute(r *gin.RouterGroup) {
	r.GET("/admin", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello from Golang in vercel")
	})
}

func init() {
	app = gin.New()
	r := app.Group("/api")
	myRoute(r)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHttp(w, r)
}