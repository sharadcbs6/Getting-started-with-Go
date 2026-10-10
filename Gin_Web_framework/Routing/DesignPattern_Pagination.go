package main

import (
	"crypto/x509"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main(){
	router:=gin.Default()
	router.GET("/api/article",func(ctx *gin.Context){
		limit,_:=strconv.Atoi(ctx.DefaultQuery("limit","20"))
		offset,_:=strconv.Atoi(ctx.DefaultQuery("offset","0"))

		if limit>100{
			limit=100
		}
		if(offset>200){
			offset=200
		}

		ctx.JSON(http.StatusOK,gin.H{
			"success": true,
			"data":	[]gin.H{
			},
			"meta": gin.H{
				"limit": limit,
				"offset": offset,
				"total": 0,
			},
		})
	})
	router.Run(":8080")
}