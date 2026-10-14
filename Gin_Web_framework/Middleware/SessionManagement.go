/*
		Context - Session Management
	sessions allows to store user-specific data	across multiple HTTP requests.
	Since HTTP is stateless, sessions use cookies to identify returning users ,and
	retrieve their stored data


*/

package main

import (
	"net/http"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main(){
	router:=gin.Default()
	// a cookie based session store using a secruity key
	store := cookie.NewStore([]byte("f|YdN;UKY!XtI.ua"))	//used randomkeygenerator for generating security key

	router.Use(sessions.Sessions("mysession",store))
	router.GET("/login",func(ctx *gin.Context) {
		session:=sessions.Default(ctx);
		session.Set("user","Sharad");
		session.Save();
		ctx.JSON(200,gin.H{"message":"logged in"})
	})
	router.GET("/profile",func(ctx *gin.Context) {
		session:=sessions.Default(ctx);
		user:=session.Get("user");
		if user==nil{
			ctx.JSON(http.StatusUnauthorized,gin.H{
				"error":"not logged in",
			})
			return
		}
			ctx.JSON(200,gin.H{"user":user})

	})
	router.GET("/logout",func(ctx *gin.Context) {
		session:=sessions.Default(ctx)
		session.Clear()
		session.Save()
		ctx.JSON(200,gin.H{"message":"logged out"})

	})
	router.Run(":8080")
}