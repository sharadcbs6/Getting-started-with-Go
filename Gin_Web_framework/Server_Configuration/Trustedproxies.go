package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)
func main(){
	router:=gin.Default();
	router.SetTrustedProxies([]string{"192.168.8.7"})
	router.GET("/",func(ctx *gin.Context) {
		fmt.Printf("ClientIP:%s\n",ctx.ClientIP());
	})
	router.Run(":8080")
}