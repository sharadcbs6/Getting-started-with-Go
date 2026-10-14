/*
				Context-Skip Logging
				You can skip logging for specific paths or based on custom logic using LoggerConfig.
				SkipPaths excludes specific routes from logging — useful for health checks or metrics endpoints that generate noise.
				Skip is a function that receives the *gin.Context and returns true to skip logging — useful for conditional logic like skipping successful responses.
*/

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	router:=gin.New()
	loggerConfig:=gin.LoggerConfig{SkipPaths: []string{"/metrics"}}
	loggerConfig.Skip=func(c *gin.Context) bool {
		return c.Writer.Status()<500;
	}

	router.Use(gin.LoggerWithConfig(loggerConfig))
	router.Use(gin.Recovery())
	router.GET("/metrics",func(ctx *gin.Context) {
		ctx.Status(http.StatusNotImplemented)
	})
	router.GET("/ping",func(ctx *gin.Context) {
		ctx.String(200,"pong");
	})
	router.GET("/data",func(ctx *gin.Context) {
		ctx.Status(http.StatusNotImplemented)
	})
	router.Run(":8080")
}