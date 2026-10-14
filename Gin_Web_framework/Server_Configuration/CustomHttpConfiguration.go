package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func main(){
	router:=gin.Default()
	router.GET("/ping",func(ctx *gin.Context) {
		ctx.String(200,"pong")
	})
	http.ListenAndServe(":8080",router)
}