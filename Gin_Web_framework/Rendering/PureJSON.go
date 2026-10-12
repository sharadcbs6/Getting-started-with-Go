/*
Context - PureJSON
JSON ingores non ASCII charater and convert it into /uXXXX non escaping characters.
pureJSON accepts non ASCII Characters and give output same as been written inside code.
*/
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.GET("/json", func(ctx *gin.Context ) {
		ctx.JSON(200, gin.H{
			"html": "<b>Hello</b>",
		})
	})
	router.GET("/purejson", func(ctx *gin.Context) {
		ctx.PureJSON(200, gin.H{
			"html": "<b>Hello</b>",
		})
	})
	router.Run(":8080")
}
