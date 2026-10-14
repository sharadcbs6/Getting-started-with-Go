package main
import(
	"github.com/gin-gonic/gin"
	"github.com/danielkov/gin-helmet/ginhelmet"
)

func main(){
	router:=gin.Default()
	router.Use(ginhelmet.Default())
	router.GET("/",func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"message":"Default Security Header used with ginHelmet",
		})
	})
	router.Run(":8080");
}