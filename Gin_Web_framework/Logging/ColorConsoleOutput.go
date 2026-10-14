package main

import "github.com/gin-gonic/gin"

func main(){
	gin.ForceConsoleColor();
	router:=gin.Default()
	router.GET("/ping",func(ctx *gin.Context) {
		ctx.String(200,"Pong");
	})
	router.Run(":8080");
}