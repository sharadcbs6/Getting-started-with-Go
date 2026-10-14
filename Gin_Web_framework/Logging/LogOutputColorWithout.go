package main

import "github.com/gin-gonic/gin"

func main(){
	gin.DisableConsoleColor();
	router:=gin.Default()
	router.GET("/ping",func(ctx *gin.Context) {
		ctx.String(200,"pong")
	})
	router.Run(":8080");
}