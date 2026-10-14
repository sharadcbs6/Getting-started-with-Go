package main
import(

	"github.com/gin-gonic/gin"
)

func main(){
	router:=gin.Default()
	expectedHost:="localhost:8080"
	router.Use(func(ctx *gin.Context) {
		if ctx.Request.Host!=expectedHost{
			ctx.AbortWithStatusJSON(400,gin.H{
				"error":	"Invalid Host request",
			})
		}
			ctx.Header("X-Frame-Options","DENY")
			ctx.Header("Content-Security-Policy","default-src:'self';connect-src *;font-src *;script-src-elem * 'unsafe-inline'	")
			ctx.Header("X-XSS-Protection","1;mode=block")
			ctx.Header("Strict-Transport-Secruity","max-age=31536000;includeSubDomains")
			ctx.Header("Referrer-Policy","strict-origin")
			ctx.Header("X-Content-Type-Options","nosniff")
			ctx.Header("Permissions-Policy","geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
			ctx.Next()	
	})
	router.GET("/ping",func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"message":"pong",
		})
	})
	router.Run()
}