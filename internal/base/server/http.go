package server

import "github.com/gin-gonic/gin"

func NewHTTPServer(debug bool) *gin.Engine {
	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) { c.String(200, "ok") })

	return r
}
