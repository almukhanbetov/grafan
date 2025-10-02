package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// healthcheck для Prometheus
	r.GET("/metrics", func(c *gin.Context) {
		c.String(http.StatusOK, "# HELP go_app up\n# TYPE go_app gauge\ngo_app 1\n")
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}

