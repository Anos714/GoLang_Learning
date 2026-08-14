package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func NewRouter() *gin.Engine {
	r := gin.Default()

	// ping route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "pong",
		})
	})

	// api versioning or grouping
	// v1:=r.Group("/api/v1"){

	// }

	return r

}
