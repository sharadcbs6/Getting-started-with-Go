package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)
func main(){
	router:=gin.Default()
	router.GET("/set-cookie",func (ctx *gin.Context)  {
		ctx.SetCookieData(&http.Cookie{
			Name:"session_id",
			Value:"abc123",
			Path:"/",
			Domain: "localhost",
			Expires: time.Now().Add(24*time.Hour),
			MaxAge: 86400,
			Secure: true,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,

		})
		ctx.String(200,"ok")
	})	
	router.Run(":8080")
}